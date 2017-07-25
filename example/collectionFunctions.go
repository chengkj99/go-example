package example

import (
	"fmt"
	"strings"
)

// 返回t在vs中第一次出现的索引，如果没有找到t，返回－1
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// 如果t存在于vs中，那么返回true，否则false
func Incloud(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// 如果使用vs中的任何一个字符串作为函数f的参数可以让f返回true，
// 那么返回true，否则false
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// 如果分别使用vs中所有的字符串作为f的参数都能让f返回true，
// 那么返回true，否则返回false
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// 返回一个新的字符串切片，切片的元素为vs中所有能够让函数f
// 返回true的元素
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// 返回一个bool类型切片，切片的元素为vs中所有字符串作为f函数
// 参数所返回的结果
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func CollectionFunc() {
	var strs = []string{"peach", "apple", "aear", "plum"}
	fmt.Println(Index(strs, "plum"))
	fmt.Println(Incloud(strs, "peach"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "a")
	}))
	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "a")
	}))
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.HasPrefix(v, "a")
	}))
}
