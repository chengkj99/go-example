package example

import "fmt"

func FunctionFunc() {
	v1 := plus(1, 2)
	v2 := plusPlus(2, 3)
	a, b := vals()
	fmt.Println("...", v1)
	fmt.Println("***", v2)
	fmt.Println("&&&", a, b)

	sum(1, 2, 3)
	sum(4, 5, 6)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b int) int {
	return a * b
}

func vals() (int, int) {
	return 100, 200
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("total:", total)
}

// 简化参数定义 a int, b int => a, b int
// 多值返回 return a, b
// 可变参数 nums ...int
