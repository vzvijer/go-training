// Package triangle implements functions to determine type of triangles
package triangle

import (
	"math"
	"sort"
)

type Kind uint

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides returns type of triangle
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	sides := []float64{a, b, c}
	sort.Float64s(sides)
	for _, side := range sides {
		if math.IsNaN(side) || math.IsInf(side, -1) || math.IsInf(side, 1) || side == 0 {
			return NaT
		}
	}
	a, b, c = sides[0], sides[1], sides[2]
	if a+b < c {
		k = NaT
	} else {
		if a == b && b == c {
			k = Equ
		} else if a == b || b == c || a == c {
			k = Iso
		} else {
			k = Sca
		}
	}
	return k
}
