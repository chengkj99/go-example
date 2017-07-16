package example

import (
	"fmt"
)

func NonBlockingChannelOperationsFunc() {
	// 如果没有设置容量，或者容量设置为0, 说明Channel没有缓存，只有sender和receiver都准备好了后它们的通讯(communication)才会发生(Blocking)。
	// 如果设置了缓存，就有可能不发生阻塞， 只有buffer满了后 send才会阻塞， 而只有缓存空了后receive才会阻塞。一个nil channel不会通信。
	message := make(chan string)
	signals := make(chan bool, 1)
	select {
	case msg := <-message:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case message <- msg:
		fmt.Println("received message", message)
	default:
		fmt.Println("no message received")
	}

	signals <- true
	select {
	case msg := <-message:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
