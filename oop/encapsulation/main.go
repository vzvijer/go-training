package main

import (
	"fmt"
	"log"

	"github.com/vzvijer/go_training/oop/encapsulation/person"
)

func main() {
	p, err := person.CreatePerson("Vlado Zvijer", 203, 125)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(p)
	p.SetName("Vuk Zvijer")
	fmt.Println(p)
}
