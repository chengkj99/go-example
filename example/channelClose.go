package example

import (
	"fmt"
)

// 关闭通道的意思是该通道将不再允许写入数据。
// 这个方法可以让通道数据的接收知道数据已经全部发送完成了。
func ChannelCloseFunc() {

	jobs := make(chan int, 3)
	done := make(chan bool)
	go func() {
		for {
			j, more := <-jobs
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
	// 因为上面协程中有个for循环，我们必须在写入三个值后关闭协程，
	// 这样接收方在遍历后能够停止，不然会一直遍历等待值的到来，以至于锁住。
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}
