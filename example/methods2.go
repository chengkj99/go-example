package example

import "fmt"

// 从某种意义上说，方法是函数的“语法糖”。当函数与某个特定的类型绑定，
// 那么它就是一个方法。也证因为如此，我们可以将方法“还原”成函数。
func MethodsFunc2() {

	p := Person{1, "cheng"}
	p.test(1)

	var f1 func(int) = p.test
	f1(2)

	Person.test(p, 3)

	var f2 func(Person, int) = Person.test
	f2(p, 4)
}

type Person struct {
	Id   int
	Name string
}

func (this Person) test(x int) {
	fmt.Println("Id:", this.Id, "Name:", this.Name)
	fmt.Println("x=", x)
}
