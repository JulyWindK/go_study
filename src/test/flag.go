package test

import (
	"flag"
	"fmt"
)

func flagTest() {
	var name string
	var age int64
	var married bool
	flag.StringVar(&name, "name", "root", "user name")
	flag.Int64Var(&age, "age", 25, "user age")
	flag.BoolVar(&married, "married", false, "user is or not married")

	flag.Parse()

	fmt.Println(name, age, married)

	fmt.Println("the rest args:", flag.Args())
	fmt.Println("the rest arg nums:", flag.NArg())
	fmt.Println("the total args of use:", flag.NFlag())
}

func flagSetTest() {
	var nameflag string
	var ageflag int
	var marriedflag bool

	// bool类型的flag 有这个flag就算true 没有就是默认值
	args := []string{"-name", "kfx", "-age", "28"}

	fs := flag.NewFlagSet("MyFlagSet", flag.ContinueOnError)
	fs.StringVar(&nameflag, "name", "root", "user name")
	fs.IntVar(&ageflag, "age", 0, "user age")
	fs.BoolVar(&marriedflag, "married", false, "user is or not married")

	fs.Parse(args)

	fmt.Println("name:", nameflag)
	fmt.Println("age:", ageflag)
	fmt.Println("married:", marriedflag)
}

func FlagTest() {
	// flagTest()
	flagSetTest()
}
