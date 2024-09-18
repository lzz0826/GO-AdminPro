package impTest

import "sync"

// 定义了一个全局的接口变量 SomeServer，用来存放实现 SomeInterFace 接口的具体对象
var SomeServer SomeInterFace

// sync.Once 用来确保某个操作只执行一次，主要用于单例模式的实现
var once sync.Once

// SomeInterFace 接口定义，它有一个方法 DoSomeThing，接受一个整型参数并返回一个整型结果
type SomeInterFace interface {
	DoSomeThing(somInt int) int
}

// SomeImp 是实现 SomeInterFace 接口的结构体类型
type SomeImp struct {
	Name string
}

func (s *SomeImp) DoSomeThing(someInt int) int {
	s.Name = "tony"    // 修改 Name 字段
	return someInt + 1 // 返回传入整数的加 1 结果
}

// init 函数会在包初始化时自动执行，它使用 sync.Once 确保 SomeImp 对象只被初始化一次
func init() {
	once.Do(func() {
		SomeServer = new(SomeImp) // 使用 new() 创建 SomeImp 的实例并赋值给 SomeServer
	})
}
