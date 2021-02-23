package algorithms

import (
	"errors"
)

func DecToBase(n, b int) (string, error) {
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}

	if n < 0 || b < 0 {
		return "", errors.New("Number and base must be positive values")
	}
	if b > 16 {
		return "", errors.New("Cannot do conversion for base greater then 16")
	}

	res := ""
	for n > 0 {
		res = digits[n%b] + res
		n = n / b
	}
	return res, nil
}
