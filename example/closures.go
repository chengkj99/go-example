package example

import "fmt"

func ClosuresFunc() {

	nextInt := intSeq()
	fmt.Print(".", nextInt())
	fmt.Print("..", nextInt())
	fmt.Print("...", nextInt())
	fmt.Println("....", nextInt())

	newInt := intSeq()
	fmt.Print(".", newInt())
	fmt.Println("..", newInt())
}

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
