package example

import "fmt"

type person struct {
	name string
	age  int
}

func StructsFunc() {
	// 定义结构体变量
	fmt.Println(".", person{"cheng", 18})
	fmt.Println("..", person{name: "kang", age: 18})
	fmt.Println("...", person{name: "jian"})

	s := &person{name: "Ann", age: 20}
	fmt.Println("~~", s)

	z := person{name: "zhonghua", age: 22}
	fmt.Println("!!", z.name, z.age)
	z.age = 26
	fmt.Println("!!", z.name, z.age)
	sp := &z
	fmt.Println("!!!", sp.name)

	sp.age = 33
	fmt.Println("!!!!", sp)

}

//定义： type name struct {
// 	property1 type
// 	property2 type
// }

//赋值
// name{property1：value1, property2: value2}
// name{property1, property2}

// 获取
// name.property

// 指针
// &name{}
