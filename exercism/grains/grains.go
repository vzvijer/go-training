package grains

import (
	"errors"
	"math"
)

func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("Invalid n")
	}

	return 1 << (n - 1), nil
}

func Total() uint64 {
	return math.MaxUint64
}
