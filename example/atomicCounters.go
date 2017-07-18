package example

import (
	"fmt"
	"sync/atomic"
	"time"
)

// Go里面的管理协程状态的主要机制就是通道通讯。 // ? 4
// 通讯有两种方式： channel， 地址传值
// 这些我们上面的例子介绍过。这里还有一些管理状态的机制，下面我们看看多协程原子访问计数器的例子，
// 这个功能是由sync/atomic包提供的函数来实现的。
func AtomicCountersFunc() {
	var ops uint64 = 0
	for i := 0; i < 2; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)    //? 1
				time.Sleep(time.Millisecond) //? 2
			}
		}()
	}

	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal, ops)
	// atomic.LoadUint64(&ops)和ops区别很小，一个是原子取值，一个是直接取值
}
