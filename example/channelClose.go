package example

import (
	"fmt"
)

// 关闭通道的意思是该通道将不再允许写入数据。
// 这个方法可以让通道数据的接收知道数据已经全部发送完成了。
func ChannelCloseFunc() {

	jobs := make(chan int, 8)
	done := make(chan bool)
	go func() {
		for {
			j, more := <-jobs
			fmt.Println("!!!", j, more)
			if more {
				fmt.Println("received jobs", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}

	}()

	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Println("send job", i)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}
