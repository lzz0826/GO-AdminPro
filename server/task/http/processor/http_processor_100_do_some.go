package processor

import (
	"AdminPro/server/task/http/request"
	"AdminPro/server/task/submit"
	"AdminPro/server/task/submit/tasks"
	"fmt"
)

func Http_Processor_100_Do_Some(p *request.HttpTaskRequest) bool {

	fmt.Printf("set Http_Processor_100_Do_Some")

	parameter := p.RequestParameter

	if roomId, exRoomId := parameter["room_id"]; exRoomId {
		fmt.Printf("roomId : ", roomId)

	} else {
		fmt.Printf("NO roomId : ", roomId)

	}

	roomId2 := parameter["room_id"].(float64)

	requsetCode := p.RequestCode
	task := tasks.NewTask(requsetCode, p.RequestParameter, p, int(roomId2))
	// 提交任务 Submit MAP
	err := submit.Submit(task)
	if err != nil {
		return false
	}
	return true
}
