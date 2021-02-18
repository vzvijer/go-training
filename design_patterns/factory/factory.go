package main

import "fmt"

type Vehicle interface {
	Drive()
}

type car struct {
	regNumber string
}

type bus struct {
	regNumber string
}

func (c car) Drive() {
	fmt.Printf("Car %s -> Brrrrrrrm", c.regNumber)
	fmt.Println()
}

func (b bus) Drive() {
	fmt.Printf("Bus %s -> BRRRRBRRRRRRRRRRRRM", b.regNumber)
	fmt.Println()
}

//factory function form making a new car
func NewCar(regNum string) Vehicle {
	return &car{regNum}
}

//factory function form making a new bus
func NewBus(regNum string) Vehicle {
	return &bus{regNum}
}

//factory function that returns car or bus depending on weight
func NewVehicle(regNum string, weight int) Vehicle {
	if weight > 2000 {
		return NewBus(regNum)
	} else {
		return NewCar(regNum)
	}
}

func testFactory() {
	vehicles := []Vehicle{
		NewBus("as123456"),
		NewCar("qw321654"),
		NewVehicle("zx678900", 2500),
		NewVehicle("zx678901", 3500),
		NewVehicle("zx678902", 500),
	}
	for _, vehicle := range vehicles {
		vehicle.Drive()
	}
}
