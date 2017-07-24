package example

import "os"

// Panic表示的意思就是有些意想不到的错误发生了。
// 通常我们用来表示程序正常运行过程中不应该出现的，或者我们没有处理好的错误。
func GoPanicFunc() {
	// 我们使用panic来检查预期不到的错误
	panic("a problem")

	// Panic的通常使用方法就是如果一个函数
	// 返回一个我们不知道怎么处理的错误的时候，直接终止执行。
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
