package example

import (
	"fmt"
	"time"
)

// 我们将在worker函数里面运行几个并行实例
// 这个函数从jobs通道里面接受任务，然后把运行结果发送到results通道。
// 每个job我们都休眠一会儿，来模拟一个耗时任务。
func WorkerPoolsFunc() {
	// 为了使用我们的工作池，我们需要发送工作和接受工作的结果，
	// 这里我们定义两个通道，一个jobs，一个results
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 这里启动3个worker协程，一开始的时候worker阻塞执行，因为
	// jobs通道里面还没有工作任务
	for w := 1; w < 3; w++ {
		fmt.Println("...", w)
		go workers(w, jobs, results)
	}
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)
	// 然后我们从results里面获得结果
	for a := 1; a <= 9; a++ {
		<-results
	}

}

// 一个工作池 FIFO 符合队列先进先出的特性

func workers(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job ", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}
