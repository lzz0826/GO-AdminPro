package tasks

import (
	"AdminPro/common/utils"
	"AdminPro/server/task/http/request"
	"context"
	"sync/atomic"
)

// TODO
type Task struct {
	Id          string                   // id
	TaskId      int                      // 请求号
	IsValid     bool                     // 任务是否有效 0无效 1有效
	DelayTime   int                      // 任务延迟执行时间(单位ms)
	BeginTime   int64                    // 任务开始时间
	LeftTime    float64                  // 任务还剩多少时间开始执行
	Map         map[string]interface{}   // 请求参数
	BaseRequest *request.HttpTaskRequest // 请求头讯息

	RoomId int //自订需求

	TaskFuture interface{} // 底成任务语句
	Cancel     atomic.Uint32
	context.Context
}

func NewTask(taskId int, m map[string]interface{}, request *request.HttpTaskRequest, roomId int) *Task {
	return &Task{
		TaskId:      taskId,
		Map:         m,
		RoomId:      roomId,
		BaseRequest: request,
		Id:          utils.GenerateUUID(),
		Context:     context.Background(),
		IsValid:     true,
	}
}

func NewTask2(taskId int, m map[string]any, request *request.HttpTaskRequest, roomId int, leftTime float64) *Task {
	return &Task{
		TaskId:      taskId,
		Map:         m,
		RoomId:      roomId,
		BaseRequest: request,
		Id:          utils.GenerateUUID(),
		LeftTime:    leftTime,
		Context:     context.Background(),
		IsValid:     true,
	}
}
