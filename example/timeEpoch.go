package example

import (
	"fmt"
	"time"
)

func EpochFunc() {
	p := fmt.Println

	//  使用Unix和UnixNano来分别获取从Unix起始时间
	// 到现在所经过的秒数和微秒数
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()

	// 注意这里没有UnixMillis方法，所以我们需要将
	// 微秒手动除以一个数值来获取毫秒
	millis := nanos / 1000000
	p(secs)
	p(millis)
	p(nanos)

	// 反过来，你也可以将一个整数秒数或者微秒数转换为对应的时间

	p(time.Unix(secs, 0))
	p(time.Unix(0, nanos))
}
