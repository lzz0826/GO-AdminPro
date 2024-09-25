package lockTest

import (
	"AdminPro/common/utils"
	"fmt"
	"sync"
	"time"
)

// 使用 ReentrantMutex 的对象
type MyObject struct {
	mutex *utils.ReentrantMutex // 使用可重入锁
	data  int                   // 模拟一些共享资源
}

// 构造函数，初始化 MyObject，并使用可重入锁
func NewMyObject() *MyObject {
	return &MyObject{
		mutex: &utils.ReentrantMutex{}, // 初始化可重入锁
		data:  0,
	}
}

// 模拟第一个操作，这里可以被重入
func (obj *MyObject) Method1(id int) {
	fmt.Printf("Goroutine %d:  Method1: Trying to lock\n", id)
	obj.mutex.Lock() // 获取可重入锁
	defer obj.mutex.Unlock()

	fmt.Printf("Goroutine %d:  Method1: Locked\n", id)
	obj.data += 10 // 模拟修改共享资源
	time.Sleep(1 * time.Second)
	obj.Method2(id) // 调用另一个方法，实现重入

	fmt.Printf("Goroutine %d:  Method1: Finished\n", id)
}

// 模拟第二个操作，可以被重入
func (obj *MyObject) Method2(id int) {
	fmt.Printf("Goroutine %d:  Method2: Trying to lock\n", id)
	obj.mutex.Lock() // 获取可重入锁
	defer obj.mutex.Unlock()

	fmt.Printf("Goroutine %d:  Method2: Locked\n", id)
	obj.data += 5 // 模拟修改共享资源
	time.Sleep(500 * time.Millisecond)

	fmt.Printf("Goroutine %d:  Method2: Finished\n", id)
}

// 模拟多个 Goroutine 并发访问同一对象
func TestMyObjectWithReentrantLock() {
	obj := NewMyObject() // 创建对象，使用可重入锁
	var wg sync.WaitGroup

	// 启动多个 Goroutine
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			fmt.Printf("Goroutine %d: Calling Method1\n", id)
			obj.Method1(id) // 每个 Goroutine 调用 Method1，进入锁
			fmt.Printf("Goroutine %d: Method1 call finished\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("All Goroutines have finished.")
}
