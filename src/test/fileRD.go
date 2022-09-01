package test

import "os"

func fileRDTest() {
	_, _ = os.Open("./defer.go")

}
