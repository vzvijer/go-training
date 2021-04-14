package main

import "flag"

func main() {
	flag.String("a", "a", "a is a flag")
	flag.Parse()

}
