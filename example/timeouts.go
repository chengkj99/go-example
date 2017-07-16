package example

import (
	"fmt"
	"time"
)

// 超时对那些连接外部资源的程序来说是很重要时，就需要限定执行时间。
// 在Go里面实现超时。很简单，我们可以使用channel和select容易地做到。
func TimeoutsFunc() {

	// 在这个例子中，假设我们执行了一个外部调用，2秒之后将结果写入c1
	c1 := make(chan int, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- 1
	}()
	// 这里使用select来实现超时。
	// `res := <-c1`等待通道结果，
	// `<- Time.After`则在等待1秒后执行。
	// select首先执行那些不再阻塞的case，
	// 如果`res := <-c1`超过1秒没有执行的话，这里会执行超时程序，
	select {
	case res := <-c1:
		fmt.Println("res one", res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1s")
	}
	// 如果将超时时间设置到3s,这个时候`res := <-c2`将在
	// 超时case之前执行，从而能够输出写入通道c2的值
	c2 := make(chan int, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- 2
	}()
	select {
	case res := <-c2:
		fmt.Println("res two", res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 3s")
	}
}
