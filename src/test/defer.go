package test

import "fmt"

func DeferTest() {
	defer func() {
		fmt.Println("defer start")
	}()

	fmt.Println("start")
	fmt.Println("end")
	defer fmt.Println("defer middle")
	defer fmt.Println("defer end")

}
