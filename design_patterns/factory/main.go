package main

import "fmt"

func main() {
	fmt.Println("Testing factory pattern")
	fmt.Println("=======================")
	testFactory()
	fmt.Println()

	fmt.Println("Testing factory generator pattern")
	fmt.Println("=================================")
	testFactoryGenerator()
	fmt.Println()
}
