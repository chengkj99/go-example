package example

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// 我们可以通过原子操作管理简单的计数器状态
// 对于更复杂的状态,我们可以使用一个互斥对象跨多个了goroutine安全地进行数据访问
func MutexesFunc() {
	// 这个例子的状态就是一个map
	var state = make(map[int]int)
	// 这个`mutex`将同步对状态的访问
	var mutex = &sync.Mutex{}

	// ops将对状态的操作进行计数
	var readOps uint64 = 0
	var writeOps uint64 = 0
	total := 0

	// 这里我们启动100个协程来不断地读取这个状态
	for i := 0; i < 100; i++ {
		go func() {
			for {
				// 对于每次读取，我们选取一个key来访问，
				// mutex的`Lock`函数用来保证对状态的唯一性访问，
				// 访问结束后，使用`Unlock`来解锁，然后增加ops计数器
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				// 为了保证这个协程不会让调度器处于饥饿状态，
				// 我们显式地使用`runtime.Gosched`释放了资源控制权，
				// 这种控制权会在通道操作结束或者time.Sleep结束后自动释放。
				// 但是这里我们需要手动地释放资源控制权

				// time.Sleep(time.Millisecond) // ??: 结束后释放？区别？为什么操作次数差距很大？
				runtime.Gosched()
			}
		}()
	}

	// 同样我们使用10个协程来模拟写状态
	for w := 0; w < 10; w++ {
		go func() {
			key := rand.Intn(5)
			val := rand.Intn(100)

			mutex.Lock()
			state[key] = val
			mutex.Unlock()

			atomic.AddUint64(&writeOps, 1)
			// time.Sleep(time.Millisecond)
			runtime.Gosched()
		}()
	}

	// 主协程Sleep，让那10个协程能够运行一段时间
	time.Sleep(time.Second)

	// 输出操作次数 ??: 操作次数作用是啥
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// 最后锁定并输出状态
	mutex.Lock()
	fmt.Println("state:", state, total)
	mutex.Unlock()

}
