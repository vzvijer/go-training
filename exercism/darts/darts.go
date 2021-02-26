package darts

import (
	"math"
)

func Score(x, y float64) int {
	score := 0

	dist := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	if dist <= 1.0 {
		score = 10
	} else if dist <= 5.0 {
		score = 5
	} else if dist <= 10.0 {
		score = 1
	}

	return score
}
