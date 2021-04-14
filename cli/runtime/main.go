package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go root:", runtime.GOROOT())
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Number of CPUs:", runtime.Compiler)
	fmt.Println("Go verion:", runtime.Version())
}
