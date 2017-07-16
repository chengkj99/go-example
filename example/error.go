package example

import "fmt"
import "errors"

func ErrorFunc() {
	for _, v := range []int{7, 42} {
		if r, e := f1(v); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, v := range []int{7, 42} {
		if r, e := f2(v); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// 如果需要使用错误信息数据，需要使用类型value，ok ：= element.(T)断言获取
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}

// Go语言里面约定错误代码是函数的最后一个返回值，
// 并且类型是error，这是一个内置的接口
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 10, nil
}

// 也可以通过实现error接口的方法Error()来自定义错误
// 下面我们自定义一个错误类型来表示上面例子中的参数错误

type argError struct {
	arg  int
	prob string
}

func (self argError) Error() string {
	return fmt.Sprintln(self.arg, self.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg, nil
}

// error说一个接口类型
// go语言里约定错误代码是函数最后一个值
// 我们可以通过errors.New("生成一个error信息")

// 也可以重新实现一个新的error接口

// type myError interface {
// 	code int
// 	msg string
// }

// 此时我们可以通过下面这个方式重新定义个性的错误信息
// &myError{cade, msg}
