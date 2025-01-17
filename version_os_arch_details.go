package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello, Yuval!")
	fmt.Println("Go Version:", runtime.Version())
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Archticture:", runtime.GOARCH)
}
