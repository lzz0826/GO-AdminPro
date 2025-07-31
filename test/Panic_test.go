package test

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
	"time"
)

// 示例 1：函数自己 panic，自己 recover
func selfRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[selfRecover] 捕获 panic：", r)
		}
	}()

	panic("selfRecover 出错了")
}

// 示例 2：下层函数 panic，由上层函数 recover
func childPanic() {
	panic("childPanic 出错了")
}

func parentRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[parentRecover] 捕获 panic：", r)
		}
	}()

	childPanic()
}

// 示例 3：goroutine 中 panic，主函数无法 recover
func goroutinePanic() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("[goroutine内部] 成功捕获 panic：", r)
			}
		}()

		panic("goroutine 出错了")
	}()
}

// 示例 4：goroutine panic，外部没有 recover（会直接 crash）
func goroutineNoRecover() {
	go func() {
		// 没有 recover，panic 会直接崩溃整个程序（如果主程序没等到它）
		panic("goroutine 无 recover")
	}()
}

func TestPanic(t *testing.T) {
	fmt.Println("------ 示例 1: selfRecover ------")
	selfRecover()

	fmt.Println("------ 示例 2: parentRecover ------")
	parentRecover()

	fmt.Println("------ 示例 3: goroutine 有内部 recover ------")
	goroutinePanic()

	fmt.Println("------ 示例 4: goroutine 无 recover (可能 crash) ------")
	// 取消下面这一行注释，会看到程序直接崩溃（因为 goroutine 中 panic 没被捕获）
	// goroutineNoRecover()

	// 为了让 goroutine 有机会输出日志
	time.Sleep(1 * time.Second)

	fmt.Println("------ main 函数结束 ------")
}

func f(data string) {
	defer func() {
		if p := recover(); p != nil {
			stackBytes := make([]byte, 1024)
			stackBytes = stackBytes[:runtime.Stack(stackBytes, false)]
			errInfo := strings.ReplaceAll(string(stackBytes), "\n", "")
			fmt.Println("BigDataEncryptFieldHandler panic recover:%v stack:%s", p, errInfo)
		}
	}()
	//这里执行业务逻辑会被上面 defer panic捕获
	fmt.Println("Test data:%v", data)

}
