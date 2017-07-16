package example

import "fmt"

func ArraysFunc() {
	var a [5]int
	fmt.Println("...", a)

	a[2] = 100
	fmt.Println("set", a)
	fmt.Println("get:", a[2])

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("!!!", b)

	var aa [3][2]int
	for x := 0; x < 3; x++ {
		for j := 0; j < 2; j++ {
			aa[x][j] = x + j
		}
	}
	fmt.Println("aa~~~", aa)
	fmt.Println("aa~~~", len(aa))

}
