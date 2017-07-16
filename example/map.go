package example

import "fmt"

func MapFunc() {
	// create
	m := make(map[string]int)

	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	fmt.Println("...", m, len(m))

	delete(m, "a")

	k1, p1 := m["a"]
	k2, p2 := m["b"]
	fmt.Println("~~~", k1, p1)
	fmt.Println("~~~", k2, p2)

	nm := map[int]string{0: "a", 1: "b"}
	fmt.Println("---", nm)
}

// 创建map
//  name := make(map[keyType]valueType)
//  name := map[keyType]valueType{key: value}

// delete(name[keyValue])

// value, isExist := name[keyValue]
