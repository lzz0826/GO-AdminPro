package _map

import (
	"AdminPro/server/task/taskEnum"
	"AdminPro/server/task/tcp/processor"
	"AdminPro/server/task/tcp/request"
)

// 任务key = RequestCode
// 声明抽象方法
type tcpHandlerMethod func(tcp *request.RequestTCP) []byte

// key:RequestCode value: 以实现的tcpHandlerMethod
var TcpHandlerMapping = map[int]tcpHandlerMethod{
	taskEnum.TCP_PROCESSOR_1_DO_SOME: processor.Tcp_Processor_1_Do_Some,
}
