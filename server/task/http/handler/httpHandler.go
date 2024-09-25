package handler

import (
	"AdminPro/common/enum"
	"AdminPro/internal/glog"
	"AdminPro/server/controller"
	_map "AdminPro/server/task/http/map"
	"AdminPro/server/task/http/request"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TaskHttpHandler(c *gin.Context) {
	var req request.HttpTaskRequest
	// 将POST请求体解析到 req 结构体中
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	// 后续处理逻辑
	CheckoutOnlineUser(c)

	// 检查是否有该请求 任务
	if handler, existHandler := _map.HttpHandlerMapping[req.RequestCode]; existHandler {

		glog.Infof("Handler executed, rep value: %v", handler)

		func() {
			defer func() {
				//recover() 函数用于捕获程序中的 panic，以防止程序崩溃。
				if r := recover(); r != nil {
					glog.Errorf("handle recover err:%v", r)
				}
			}()
			// 这里可以放置可能会引发 panic 的代码
			fmt.Printf("Handler executed, rep value")

			//使用该任务 带入RequertHttp
			rep := handler(&req)
			if rep {
				controller.WebResp(c, enum.SUCCESS, true, enum.GetResponseMsg(enum.SUCCESS))
			}
		}()
	} else {
		msg := fmt.Sprintf("handle client request error, can't find requertHttp code=%v, userid=%v", req, req.UserId)
		glog.Errorf(msg)
		controller.WebResp(c, enum.SUCCESS, msg, enum.GetResponseMsg(enum.ERROR))

	}

}

func CheckoutOnlineUser(c *gin.Context) {

	//TODO Redis
	id := controller.GetCurrentAdminId(c)
	glog.Infof("OnlineUser : %d", id)

}
