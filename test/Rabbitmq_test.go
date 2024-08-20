package test

import (
	"AdminPro/common/mq"
	"fmt"
	"testing"
)

func TestRabbitmqInitLoseBetslipQueue(t *testing.T) {
	fmt.Println("----------------------------")
	mq.InitLoseBetslipQueue()
	fmt.Println("TestRabbitmq Start-")
	fmt.Println("----------------------------")
	fmt.Println("TestRabbitmq SendLoseBetSlip -")

	mq.SendLoseBetSlip("test22222")

}
