package test

import (
	"fmt"
	"net/http"
)

func HttpTest() {
	resp, _ := http.Get("http://www.google.com/robots.txt")

	fmt.Println(resp)
}
