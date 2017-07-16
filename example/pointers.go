package example

import "fmt"

func PointersFunc() {
	i := 1
	fmt.Println("initial:", i) //1

	zeroval(i)
	fmt.Println("zeroval:", i) //1

	zeroptr(&i)
	fmt.Println("zeroptr:", i) //0

	fmt.Println("pointer:", &i) //xxx

}

// 参数是变量的值
func zeroval(ival int) {
	ival = 0
}

// 参数是变量的地址
// 在函数内部对这个地址所指向的变量的任何修改都会反映到原来的变量上。
func zeroptr(ival *int) {
	fmt.Println("ival:", ival)
	*ival = 0
}
