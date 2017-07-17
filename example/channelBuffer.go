package example

import "fmt"

func ChannelBufferFunc() {
	chanMsg := make(chan string)
	go func() {
		chanMsg <- "hello world"
		chanMsg <- "hello china"
	}()

	msg1 := <-chanMsg
	msg2 := <-chanMsg
	fmt.Println(msg1)
	fmt.Println(msg2)
}

// buffer
// channelMessage := make(chan type, bufferNum)
