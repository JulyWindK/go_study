package test

import "fmt"

// FormatPrtTest
// 测试输出 %v %+v %#v 格式
//*
func FormatPrtTest() {
	s := stu{
		"zs",
		25,
	}

	// 只输出所有值
	fmt.Printf("s:%v\n", s) // s:{zs 25}

	// 先输出字段名，再输出该字段的值
	fmt.Printf("s:%+v\n", s) // s:{name:zs age:25}

	// 先输出结构体名字值，再输出结构体（字段名字+字段的值）
	fmt.Printf("s:%#v\n", s) // s:test.stu{name:"zs", age:25}

	fmt.Printf("str:%q", 0x4E2D)

}
