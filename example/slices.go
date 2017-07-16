package example

import "fmt"

func SlicesFunc() {
	s := make([]int, 3)
	fmt.Println("...", s)
	s[0] = 1
	s[1] = 2
	s[2] = 3
	fmt.Println("...", s)
	s = append(s, 4)
	s = append(s, 9)
	fmt.Println("...", s, len(s))

	c := make([]int, len(s))
	copy(c, s)
	fmt.Println(c)

	l := c[1:3]
	l1 := c[1:]
	l2 := c[:3]
	fmt.Println("$$$", l, l1, l2)

	t := []int{1, 2, 3, 4, 5, 6, 76, 8, 8, 9}
	fmt.Println("***", t)

	mm := make([][]int, 3)
	for f := 0; f < 3; f++ {
		len3 := f + 1
		mm[f] = make([]int, len3)
		for g := 0; g < len3; g++ {
			mm[f][g] = f + g
		}
	}
	fmt.Println("@@@", mm, len(mm))

}

//  make([]type, len,cap)
// []type{1,2,3,4}

// append()
// mySlice[1:2]
// copy(newS, oldS)
