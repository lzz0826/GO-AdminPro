package _map

import (
	"AdminPro/server/task/http/processor"
	"AdminPro/server/task/http/request"
	"AdminPro/server/task/taskEnum"
)

type httpHandlerMethod func(p *request.HttpTaskRequest) bool

// key:RequestCode value: 以实现的httpHandlerMethod
var HttpHandlerMapping = map[int]httpHandlerMethod{
	taskEnum.HTTP_PROCESSOR_100_DO_SOME: processor.Http_Processor_100_Do_Some,
}
