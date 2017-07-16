package example

import "fmt"

// 协程 是轻量级的线程
func GoroutineFunc() {
	f("direct")

	// 为了能够让这个函数以协程(goroutine)方式
	// 运行使用go f(s)
	// 这个协程将和调用它的协程并行执行
	go f("goroutine")

	// 也可以为匿名函数开启一个协程运行
	go func(msg string) {
		fmt.Println("!!!", msg)
	}("hello world")

	// 上面的协程在调用之后就异步执行了，所以程序不用等待它们执行完成
	// 就跳到这里来了，下面的Scanln用来从命令行获取一个输入，然后才
	// 让main函数结束
	// 如果没有下面的Scanln语句，程序到这里会直接退出，而上面的协程还
	// 没有来得及执行完，你将无法看到上面两个协程运行的结果
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")

}
func f(from string) {
	for i := 0; i < 5; i++ {
		fmt.Println("...", from, "-", i)
	}
}

// 启用并行执行
// go f1()
// go f2()
// 这样f1和f2就是在并行执行状态
