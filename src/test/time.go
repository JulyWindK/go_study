package test

import (
	"fmt"
	"time"
)

func TimeTest() {
	start := time.Now()
	fmt.Println(start.String())
	fmt.Println(start.After(start))
	fmt.Println(start.Local())

	time.Sleep(3)

	t := time.Since(start).Seconds()
	fmt.Printf("since start: %fl s\n", t)
}
