package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hi, there!")
	fmt.Println("OS:   ", runtime.GOOS)
	fmt.Println("ARCH: ", runtime.GOARCH)
}
