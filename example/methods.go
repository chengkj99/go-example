package example

import "fmt"

// 一般的函数定义叫做函数，定义在结构体上面的函数叫做该结构体的方法。
func MethodsFunc() {
	r1 := rect{
		width:  10,
		height: 6,
	}

	fmt.Println("...", r1.area())
	fmt.Println("!!!", r1.prim())

	r1p := &r1
	fmt.Println("...", r1p.area())
	fmt.Println("!!!", r1p.prim())

}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	a := r.width * r.height
	return a
}

// 方法的定义限定类型可以为结构体类型
// 也可以是结构体指针类型
// 结构体类型可以更改成员信息
func (r rect) prim() int {
	return 2 * (r.height + r.width)
}

// 方法是限定了类型的函数，结构体类型的函数称为结构体方法
// type structname struct{}

// func funcName(self [*]structname) valuetype {}

// structname.funcName
