package tasks

import (
	"fmt"
	"time"
)

type Processor func()

func Request_1_do_some(t *Task) Processor {
	return func() {

		id := t.Id

		fmt.Printf("Do SomeThing Start...")
		fmt.Printf("Do SomeThing Start...")

		fmt.Printf("Task ID : %d\n", id)
		fmt.Printf("RoomId : %d\n", t.RoomId)
		fmt.Printf("额外参数 : %d\n", t.Map)

		time.Sleep(60 * time.Second)

		fmt.Printf("Do SomeThing End...")

	}
}
