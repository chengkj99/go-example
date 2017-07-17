package example

import "fmt"

func ChannelRangeFunc() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for e := range queue {
		fmt.Println(e)
	}
}
