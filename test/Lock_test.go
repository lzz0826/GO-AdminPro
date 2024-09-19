package test

import (
	"AdminPro/common/utils"
	"AdminPro/server/server/lockTest"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTestReentrantLock(t *testing.T) {
	utils.TestReentrantLock()

}

func TestNewReentrantLock(t *testing.T) {
	lock := utils.NewReentrantLock()
	var wg sync.WaitGroup
	numGoroutines := 3

	// 创建多个 goroutine 来同时访问 Lock 方法
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			lock.Lock()
			fmt.Printf("Goroutine %d acquired the lock\n", id)
			// 模拟耗时操作
			time.Sleep(1000 * time.Millisecond)
			fmt.Printf("Goroutine %d releasing the lock\n", id)
			lock.Unlock()
		}(i)
	}

	// 等待所有 goroutine 完成
	wg.Wait()
}

// 模拟多个 Goroutine 竞争 ReentrantMutex 的测试
func TestReentrantMutexConcurrent(t *testing.T) {
	var wg sync.WaitGroup
	mutex := &utils.ReentrantMutex{}

	goroutineCount := 5
	wg.Add(goroutineCount)

	for i := 0; i < goroutineCount; i++ {
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d: trying to acquire the lock\n", id)

			mutex.Lock() // 尝试获取锁
			fmt.Printf("Goroutine %d: acquired the lock\n", id)

			time.Sleep(1 * time.Second) // 持有锁 1 秒，模拟工作
			fmt.Printf("Goroutine %d: releasing the lock\n", id)

			mutex.Unlock() // 释放锁
		}(i)
	}

	wg.Wait() // 等待所有 Goroutine 完成
	fmt.Println("All Goroutines have finished.")
}

// 模拟多个 Goroutine 并发访问同一对象
func TestSomeLockObj_ReentrantLockDD(t *testing.T) {
	lockTest.TestMyObjectWithReentrantLock()
}
