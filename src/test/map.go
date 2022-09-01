package test

import "fmt"

// MapTest map的声明及使用
func MapTest() {
	// 先声明再赋值
	var m1 map[string]bool
	m1 = make(map[string]bool)
	m1["zs"] = true
	m1["ls"] = false

	// 直接声明
	m2 := make(map[string]bool)
	m2["a"] = false
	m2["b"] = true

	// 声明并初始化
	m3 := map[string]bool{
		"kk": true,
		"yy": false,
	}

	// 遍历map
	for k := range m1 {
		fmt.Println("m1 key", k)
	}

	for k := range m2 {
		fmt.Println("m2 key", k)
	}

	for k := range m3 {
		fmt.Println("m3 key", k)
	}

	// 查找键值是否存在
	if v, ok := m1["zs"]; ok {
		fmt.Println("[m1]find key:", v)
	} else {
		fmt.Println("[m1]not find")
	}

}
