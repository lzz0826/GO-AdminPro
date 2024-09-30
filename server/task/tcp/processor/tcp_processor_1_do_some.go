package processor

import (
	"AdminPro/common/enum"
	"AdminPro/server/task/submit"
	"AdminPro/server/task/submit/tasks"
	"AdminPro/server/task/tcp/request"
	"AdminPro/server/task/tcp/response"
	"fmt"
)

func Tcp_Processor_1_Do_Some(p *request.RequestTCP) []byte {
	fmt.Printf("set Tcp_Processor_1_Do_Some")
	userId := p.UserId
	token := p.Token
	customId := p.CustomId
	code := p.RequestCode
	requsetCode := p.RequestCode

	m := make(map[string]any)
	m["requsetCode"] = requsetCode
	m["userId"] = userId
	m["token"] = token

	task := tasks.NewTask(int(code), m, p, int(customId))

	//配置延迟时间
	task.LeftTime = 5.00

	// 提交任务 Submit MAP
	err := submit.Submit(task)
	if err != nil {
		return []byte{}
	}

	//返回相关讯息(需要转byte):
	responseByte := &response.ResponseTCP{
		Code:     int32(enum.SUCCESS),
		CustomId: 22,
	}
	repByte := responseByte.ToBytes()
	return repByte
}
