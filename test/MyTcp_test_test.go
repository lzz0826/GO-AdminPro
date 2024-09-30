package test

import (
	"AdminPro/server/task/tcp/request"
	"AdminPro/server/task/tcp/response"
	"fmt"
	"net"
	"testing"
	"time"
)

// 通用发送和接收函数
func sendAndReceiveTCP(conn net.Conn, entries []DataEntry) {
	//go func() {
	for _, entry := range entries {
		// 打印消息
		fmt.Println(entry.Message)

		// 发送数据
		_, err := conn.Write(entry.Data)
		if err != nil {
			fmt.Println("发送数据出错:", err.Error())
			return
		}

		if entry.Action == "enter" || entry.Action == "sit" || entry.Action == "add" {
			fmt.Println()
		}

		// 接收响应
		var buffer = make([]byte, 1024*1024)
		for {
			conn.SetDeadline(time.Now().Add(time.Duration(100) * time.Second))
			n, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("数据接收时出错:", err.Error())
				return
			}
			if n > 0 {
				// 处理接收到的数据
				response := response.ResponseTCP{}
				response.Bt = buffer[:n] // 将接收到的数据赋值给 Bt
				response.Size = int32(n)
				if response.Parse() { // 调用解析方法
					// 打印解析结果
					fmt.Printf("收到数据: %s\n", string(response.Bt))
					fmt.Printf("response CustomId: %v\n", response.CustomId)
					fmt.Printf("response Code: %v\n", response.Code)
					fmt.Printf("response Size: %v\n", response.Size)
				} else {
					fmt.Println("解析响应数据失败")
				}
				break
			}
		}
		time.Sleep(time.Duration(entry.SleepTime) * time.Second)
	}
	time.Sleep(time.Duration(30) * time.Second)
}

func TestTCP1(t *testing.T) {
	conn, err := net.Dial("tcp", "0.0.0.0:8040")
	if err != nil {
		fmt.Println(err)
	}
	// 创建 RequestTCP 对象
	req := &request.RequestTCP{
		RequestCode:       11111,
		Protocol:          11111,
		UserId:            11111,
		LanguageId:        11111,
		ClientPlatform:    11111,
		ClientBuildNumber: 11111,
		CustomId:          11111,
		ProductId:         11111,
		Size:              11111,
	}
	// 将 RequestTCP 转换为字节数组
	data := req.ToBytes()
	// 打印转换后的字节数组
	fmt.Printf("Serialized Data: %v\n", data)

	// 使用 InitRequestTCP 重新解析字节数组
	newReq := &request.RequestTCP{}
	success := newReq.InitRequestTCP(conn, data)
	if success {
		fmt.Printf("Parsed Request: %+v\n", newReq)
	} else {
		fmt.Println("Failed to parse request")
	}
	sendAndReceiveTCP(conn, []DataEntry{
		{
			//使用转好的 byte
			Data:      data,
			Message:   "enterroom",
			SleepTime: 1,
			Action:    "enter",
		},
	})
}

func TestTCP2(t *testing.T) {
	conn, err := net.Dial("tcp", "0.0.0.0:8040")
	if err != nil {
		fmt.Println(err)
	}
	// 创建 RequestTCP 对象
	req := &request.RequestTCP{
		RequestCode:       1,
		Protocol:          22222,
		UserId:            22222,
		LanguageId:        22222,
		ClientPlatform:    22222,
		ClientBuildNumber: 22222,
		CustomId:          22222,
		ProductId:         22222,
		Size:              22222,
	}
	// 将 RequestTCP 转换为字节数组
	data := req.ToBytes()
	// 打印转换后的字节数组
	fmt.Printf("Serialized Data: %v\n", data)

	// 使用 InitRequestTCP 重新解析字节数组
	newReq := &request.RequestTCP{}
	success := newReq.InitRequestTCP(conn, data)
	if success {
		fmt.Printf("Parsed Request: %+v\n", newReq)
	} else {
		fmt.Println("Failed to parse request")
	}
	sendAndReceiveTCP(conn, []DataEntry{
		{
			//使用转好的 byte
			Data:      data,
			Message:   "enterroom",
			SleepTime: 1,
			Action:    "enter",
		},
	})
}

func TestTCP(t *testing.T) {
	//conn, err := net.Dial("tcp", "16.162.29.171:8051")
	conn, err := net.Dial("tcp", "0.0.0.0:8040")
	//conn, err := net.Dial("tcp", "dz-fat-room.abc9by6pt.com:8051")
	if err != nil {
		fmt.Println(err)
	}
	//RequestCode:1
	//Protocol:1
	//UserId:1000942
	//LanguageId:1
	//ClientPlatform:2
	//ClientBuildNumber:254
	//CustomId:14
	//ProductId:1002
	//Size:118
	//Token:ff1e5140f7a406fa5796c07a5ea2a72
	//ItemCount:0
	data := []byte{
		0x44, 0x5A, 0x50, 0x4B, 0x01, 0x00, 0x0F, 0x45,
		0xEE, 0x01, 0x02, 0xFE, 0x00, 0x0E, 0x03, 0xEA,
		0x00, 0x01, 0x00, 0x76, 0x00, 0x32, 0x00, 0x66,
		0x00, 0x66, 0x00, 0x31, 0x00, 0x65, 0x00, 0x35,
		0x00, 0x31, 0x00, 0x34, 0x00, 0x30, 0x00, 0x66,
		0x00, 0x37, 0x00, 0x61, 0x00, 0x34, 0x00, 0x30,
		0x00, 0x36, 0x00, 0x66, 0x00, 0x61, 0x00, 0x35,
		0x00, 0x37, 0x00, 0x39, 0x00, 0x36, 0x00, 0x63,
		0x00, 0x30, 0x00, 0x37, 0x00, 0x61, 0x00, 0x35,
		0x00, 0x65, 0x00, 0x61, 0x00, 0x32, 0x00, 0x61,
		0x00, 0x37, 0x00, 0x32, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x06, 0x3C, 0x01, 0x3D, 0x01, 0x3E,
		0x01, 0x83, 0x00, 0x04, 0x00, 0x00, 0x00, 0x3D,
		0x84, 0x00, 0x04, 0x00, 0x0E, 0x00, 0x05, 0x85,
		0x00, 0x04, 0x00, 0x00, 0x00, 0x00,
	}
	sendAndReceiveTCP(conn, []DataEntry{
		{
			Data:      data,
			Message:   "enterroom",
			SleepTime: 1,
			Action:    "enter",
		},
	})
}
