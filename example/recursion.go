package example

import "fmt"

func RecursionFunc() {
	fmt.Println("...", fact(5))
}

func fact(n int) int {
	if n == 1 {
		return 1
	}
	return n * fact(n-1)
}
