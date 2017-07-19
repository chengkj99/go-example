package example

import (
	"fmt"
	"sort"
)

// Go的sort包实现了内置数据类型和用户自定义数据类型的排序功能。我们先看看内置数据类型的排序。
func SortFunc() {
	// 这些排序方法都是针对内置数据类型的。
	// 这里的排序方法都是就地排序，也就是说排序改变了切片内容，而不是返回一个新的切片.
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings..", strs)
	s1 := sort.StringsAreSorted(strs)
	fmt.Println("sorted..", s1)

	// 对于整型的排序
	ints := []int{7, 2, 4, 6}
	sort.Ints(ints)
	fmt.Println("Ints..", ints)

	// 我们还可以检测切片是否已经排序好
	s2 := sort.IntsAreSorted(ints)
	fmt.Println("sorted..", s2)

}
