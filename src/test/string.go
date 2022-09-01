package test

import (
	"fmt"
	"strings"
)

func LowerUpperTrastTest() {
	basestr := "aBcDeFg"

	lowstr := strings.ToLower(basestr)
	upstr := strings.ToUpper(basestr)

	fmt.Println("basestr:", basestr)
	fmt.Println("lowstr:", lowstr)
	fmt.Println("lowstr:", upstr)

}

func SplitTest() {
	z := strings.SplitN("a,b,c", ",", -1)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)

	z = strings.SplitN("a,b,c", ",", 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)

	z = strings.SplitN("a,b,c", ",", 1)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)

	z = strings.SplitN("a,b,c", ",", 2)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)

	z = strings.SplitN("a,b,c", ",", 3)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)

	z = strings.SplitN("a,b,c", ",", 4)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)
}

func StringTest() {
	// LowerUpperTrastTest()
	// SplitTest()
}
