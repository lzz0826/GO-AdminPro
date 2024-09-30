package submit

import (
	"AdminPro/internal/glog"
	"AdminPro/server/task/submit/tasks"
	"AdminPro/server/task/taskEnum"
	"github.com/RussellLuo/timingwheel"
	"github.com/panjf2000/ants/v2" //goroutine 池
	"time"
)

// 创建线程池
var antsPool, _ = ants.NewPool(2000)

// 创建时间轮
var timinWheel = timingwheel.NewTimingWheel(time.Millisecond, 20)

// 启动时间轮
func init() {
	timinWheel.Start()
}

// 提交任务
func Submit(task *tasks.Task) (err error) {
	switch task.TaskId {

	//使用HTTP协议
	case taskEnum.HTTP_PROCESSOR_100_DO_SOME:
		//err = antsPool.Submit(tasks.Request_1_do_some(task))
		//延迟任务
		timinWheel.AfterFunc(time.Duration(task.LeftTime)*time.Second, func() {
			err = antsPool.Submit(tasks.Request_1_do_some(task))
		})

	case taskEnum.TCP_PROCESSOR_1_DO_SOME: //使用TCP协议
		err = antsPool.Submit(tasks.Request_1_do_some(task))

	//case mdata.TASK_END_INSURANCE://延迟任务
	//	timinWheel.AfterFunc(time.Duration(taskController.LeftTime)*time.Second, func() {
	//		err = antsPool.Submit(Handle_10010(taskController))
	//	})

	default:
		glog.Errorf("unknow taskController id: %d ", task.TaskId)
		break
	}
	if err != nil {
		glog.Errorf("submit taskController err: %v ", err)
	}
	return err
}

func CleanWorkResource() {
	timinWheel.Stop()
}

// 参数delay是毫秒(Millisecond)为单位
func SubmitDelayTask(roomId int, task *tasks.Task, delay int64) {
	timinWheel.AfterFunc(time.Duration(delay)*time.Millisecond, func() {
		err := Submit(task)
		if err != nil {
			glog.Errorf("SubmitDelayTask taskController err: %v ", err)
		}
	})
}
