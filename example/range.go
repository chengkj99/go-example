package example

import "fmt"

// RangeFunc iterates over elements in a variety of data structures
func RangeFunc() {
	// iterates Slices、Array、Map
	// nums := []int{1, 2, 3}
	// arr := [3]int{2, 3, 5}
	maps := map[string]int{"a2": 2, "a3": 3, "a5": 5}

	for index, value := range maps {
		fmt.Printf("s% -> s%\n", index, value)
	}

	// iterates string
	for i, c := range "0123" {
		fmt.Println("###", i, c)
		// i index 是索引值
		// c Unicode code points
	}

}

//  range 可以遍历数组，slices，map，字符串等类型

// ？？如何通过判断不同的类型写一个函数就能搞定
// func iteratesF(data interface{}) (int, int) {
// 	switch t := data.(type) {
// 	case bool:
// 		fmt.Println(" I am bool")
// 	case string:
// 		fmt.Println(" I am string")
// 	case int:
// 		fmt.Println(" I am int")
// 	default:
// 		fmt.Printf("Don't know type %T\n", t)
// 	}
// 	sumI, sumV := 0, 0
// 	for index, value := range data {
// 		sumI += index
// 		sumV += value
// 	}
// 	return sumI, sumV
// }
