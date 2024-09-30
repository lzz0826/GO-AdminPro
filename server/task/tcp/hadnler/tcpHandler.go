package hadnler

import (
	"AdminPro/internal/glog"
	_map "AdminPro/server/task/tcp/map"
	"AdminPro/server/task/tcp/request"
	"context"
	"fmt"
	"net"
	"time"
)

func TcpHandleConnection(parentCtx context.Context, conn net.Conn) {
	fmt.Printf("TCP接收到消息.......")
	defer func() {
		err := conn.Close() // 确保连接被关闭
		if err != nil {
			glog.Error("conn close error:%+v", err)
		}
	}()
	ctx, cancel := context.WithCancel(parentCtx)
	readWriteCh := make(chan struct{}, 100)
	defer func() {
		close(readWriteCh)
	}()
	// 前端就按照正常的间格时间发送一个心跳请求
	go func() {
		ticker := time.NewTicker(15 * time.Second)
		defer ticker.Stop()
		for {
			select {
			// 处理定时器触发的情况
			case <-ticker.C:
				// 每当定时器到期时，取消上下文，通知相关 goroutine 停止工作
				cancel() // 取消上下文，退出 goroutine
				return
			case <-ctx.Done():
				// 如果上下文被取消，退出 goroutine
				return
			case <-readWriteCh:
				// 如果收到读写信号，重置定时器，开始新的 15 秒计时
				ticker.Reset(15 * time.Second) // 如果收到读写信号，重置计时器
			}
		}
	}()

	for {
		buf := make([]byte, 8*1024*2)

		//连接会在 30 秒内没有接收到数据时自动超时
		err := conn.SetReadDeadline(time.Now().Add(time.Duration(30) * time.Second))
		if err != nil {
			glog.Errorf("Set conn read deadline error.")
			cancel()
			return
		}

		requestLength, err := conn.Read(buf)
		if err != nil {
			glog.Infof("accept ConnectionError reading message length: %v", err)
			cancel()
			return
		}

		if requestLength == 0 {
			glog.Infof("accept ConnectionError reading message length: %v", err)
			cancel()
			return
		}

		//创新的RequestTCP
		request := request.RequestTCP{}
		message := make([]byte, requestLength)
		copy(message, buf)

		readWriteCh <- struct{}{}

		request.Size = int32(requestLength)

		//解析 RequestTCP
		result := request.InitRequestTCP(conn, message)
		fmt.Println("解析 RequestTCP")
		fmt.Println(result)
		fmt.Println(request.RequestCode)
		fmt.Println(request.UserId)
		fmt.Println(request.CustomId)
		fmt.Println("解析 RequestTCP")

		//TODO 验证是否为以上线的用户
		//value, isOnlineUser := cache.OnlineUserInfo.Load(request.UserId)
		//glog.Infof("parse proto  request: %v", request)
		//
		//userInfo, typeOk := value.(*model.UserInfo)

		//排除
		//if !result || (request.RequestCode != mdata.REQ_GAME_ROOM_XITIAO &&
		//	request.RequestCode != mdata.REQ_GAME_ENTER_PENDING &&
		//	request.RequestCode != mdata.REQ_GAME_ENTER_ROOM && isOnlineUser && typeOk && userInfo.Channel.RemoteAddr() != conn.RemoteAddr()) {
		//	glog.Errorf("conn illegal err: result %v RequestCode:%v  userInfo: %+v", result, request.RequestCode, userInfo)
		//	return
		//}

		//glog.Infof("auth success:  requestCode= %d", request.RequestCode)

		// 检查是否有该请求 任务
		if handler, existHandler := _map.TcpHandlerMapping[int(request.RequestCode)]; existHandler {
			func() {
				defer func() {
					if r := recover(); r != nil {
						glog.Errorf("handle recover err:%v", r)
					}
				}()
				//使用该任务 带入RequestTCP
				responseByte := handler(&request)
				if len(responseByte) > 0 {
					err := conn.SetWriteDeadline(time.Now().Add(time.Duration(30) * time.Second))
					if err != nil {
						glog.Infof("conn write err :%v", err)
						return
					}
					writeLength, writeErr := conn.Write(responseByte)
					if writeErr != nil {
						glog.Errorf("conn write err :%v", writeErr)
						return
					}
					glog.Infof("conn write  success writeLength ?%v", writeLength)
				}
			}()
		} else {
			glog.Errorf("handle client request error, can't find request code=%v, userid=%v", request.RequestCode, request.UserId)
		}
	}
}
