package test

import "fmt"

type people interface {
	say(words string)
}

type stu struct {
	name string
	age  int
}

func (s *stu) say(words string) {
	fmt.Println(words)
}

func NewStu(name string, age int) people {
	s := &stu{
		name,
		age,
	}
	return s
}
