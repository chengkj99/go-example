package example

import (
	"fmt"
)

func ChannelFunc() {
	// 使用`make(chan 数据类型)`来创建一个Channel
	// Channel的类型就是它们所传递的数据的类型
	message := make(chan string)

	// 使用`channel <-`语法来向一个Channel写入数据
	// 这里我们从一个新的协程向messages通道写入数据ping
	go func() {
		message <- "hello world"
	}()

	// 使用`<-channel`语法来从Channel读取数据
	// 这里我们从main函数所在的协程来读取刚刚写入
	// messages通道的数据
	msg := <-message
	fmt.Println("...", msg)

	// 当我们运行程序的时候，数据ping成功地从一个协程传递到了另外一个协程。
	// 默认情况下，协程之间的通信是同步的，也就是说数据的发送端和接收端必须配对使用。
	// Channel的这种特点使得我们可以不用在程序结尾添加额外的代码也能够获取协程发送端发来的信息。
	// 因为程序执行到msg:=<-messages的时候被阻塞了，直到获得发送端发来的信息才继续执行。
}

// 并行通道 channel

// 创建
// channelName := make(chan type)

// 写入数据
// go func(){
// 	channelName <- "messgae"
// }()

// 获取数据
// getMsg := <-channelName
