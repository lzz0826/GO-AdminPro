package taskTest

import (
	"AdminPro/common/utils"
	"context"
	"sync/atomic"
)

// TODO
type Task struct {
	Id          string              // id
	TaskId      int                 // 请求号
	IsValid     bool                // 任务是否有效 0无效 1有效
	DelayTime   int                 // 任务延迟执行时间(单位ms)
	BeginTime   int64               // 任务开始时间
	LeftTime    float64             // 任务还剩多少时间开始执行
	Map         map[int]interface{} // 请求参数字典
	BaseRequest *Request            // 请求头讯息

	RoomId   int //自订需求
	RoomPath int //自订需求

	TaskFuture interface{} // 底成任务语句
	Cancel     atomic.Uint32
	context.Context
}

func NewTask(taskId int, m map[int]any, request *Request, roomId, roomPath int) *Task {
	return &Task{
		TaskId:      taskId,
		Map:         m,
		RoomId:      roomId,
		RoomPath:    roomPath,
		BaseRequest: request,
		Id:          utils.GenerateUUID(),
		Context:     context.Background(),
		IsValid:     true,
	}
}

func NewTask2(taskId int, m map[int]any, request *Request, roomId, roomPath int, leftTime float64) *Task {
	return &Task{
		TaskId:      taskId,
		Map:         m,
		RoomId:      roomId,
		RoomPath:    roomPath,
		BaseRequest: request,
		Id:          utils.GenerateUUID(),
		LeftTime:    leftTime,
		Context:     context.Background(),
		IsValid:     true,
	}
}
