package test

import "fmt"

// 测试切片打散传递/*
func sliceTest(s ...stu) {
	fmt.Printf("len: %d\n", len(s))
	for i, v := range s {
		fmt.Printf("[%d] %+v\n", i, v)
	}
}

func SliceTest() {
	s := []stu{{"wmz", 30}, {"sss", 48}}
	sliceTest(s...)
}
