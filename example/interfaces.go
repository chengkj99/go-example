package example

import (
	"fmt"
	"math"
)

// 接口是方法命名集合
func InterfaceFunc() {
	// 这里circle和square都实现了geometry接口，所以
	// circle类型变量和square类型变量都可以作为measure函数的参数
	s := square{10, 10}
	c := circle{10}
	measure(s)
	measure(c)

}

type geometry interface {
	area() float64
	prime() float64
}

type square struct {
	width, height float64
}
type circle struct {
	radius float64
}

// 在去实现一个接口，我们只需要实现一个接口所有的方法
// 正方形实现的接口
func (self square) area() float64 {
	return self.width * self.height
}
func (self square) prime() float64 {
	return (2 * (self.width + self.height))
}

// 圆形实现的接口
func (self circle) area() float64 {
	return math.Pi * self.radius * self.radius
}
func (self circle) prime() float64 {
	return 2 * math.Pi * self.radius
}

// 如果一个变量是接口类型，那么我们可以通过命名接口变量调用接口下的方法
// 这是一个集合图形的公共测量函数
func measure(g geometry) {
	fmt.Println("...", g)
	fmt.Println("...", g.area())
	fmt.Println("...", g.prime())
	fmt.Println("------")
}

// 定义接口
// type interfaceName interface {
// 	methodsName type
// }

// 定义结构体
// type structName struct {
// 	property type
// }

// 实现接口
// func (self structName) methodsName() type {...}

// 定义使用接口公共函数
// func funcName(self interfaceName) {...}

//使用函数
// structNameExample := structName{"propertyName"}
// structName实现了interfaceName接口 所以可以使用
// funcName(structNameExample)
