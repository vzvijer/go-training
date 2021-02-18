package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

func (e Employee) String() string {
	return fmt.Sprintf("Employee: %s, %s, %d", e.Name, e.Position, e.AnnualIncome)
}

func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{
			Name:         name,
			Position:     position,
			AnnualIncome: annualIncome,
		}
	}
}

func testFactoryGenerator() {
	employees := make([]Employee, 0)
	developerFactory := NewEmployeeFactory("Developer", 500000)
	managerFactory := NewEmployeeFactory("Manager", 800000)
	employees = append(employees, *developerFactory("Alan"))
	employees = append(employees, *managerFactory("Bob"))
	employees = append(employees, *developerFactory("Chris"))

	for _, employee := range employees {
		fmt.Println(employee)
	}
}
