package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Runtime GOOS: ", runtime.GOOS)
	fmt.Println("Runtime ARCH: ", runtime.GOARCH)
}
