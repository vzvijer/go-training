package person

import (
	"errors"
	"fmt"
)

type Person struct {
	name   string
	height int
	weight int
}

func CreatePerson(name string, height, weight int) (*Person, error) {
	if height <= 0 {
		return nil, errors.New("Height of a person can't be less or equal to 0!")
	}
	if weight <= 0 {
		return nil, errors.New("Height of a person can't be less or equal to 0!")
	}
	return &Person{
		name:   name,
		height: height,
		weight: weight,
	}, nil
}

func (p Person) String() string {
	return fmt.Sprintf("name: %s, height: %d, weight: %d", p.name, p.height, p.weight)
}

func (p *Person) Name() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}
