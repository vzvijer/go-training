package perfect

import (
	"errors"
)

type Classification uint

var ErrOnlyPositive error = errors.New("Accepting only natural numbers.")

const (
	ClassificationUnknown = iota
	ClassificationPerfect
	ClassificationAbundant
	ClassificationDeficient
)

func factors(n int64) []int64 {
	factors := []int64{}
	for i := int64(1); i <= n/2; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

func sum(s []int64) int64 {
	sum := int64(0)
	for _, n := range s {
		sum += n
	}
	return sum
}

func Classify(n int64) (Classification, error) {
	var c Classification
	if n <= 0 {
		return 0, ErrOnlyPositive
	}
	sum := sum(factors(n))
	switch {
	case n == sum:
		c = ClassificationPerfect
	case n < sum:
		c = ClassificationAbundant
	case n > sum:
		c = ClassificationDeficient
	default:
		c = ClassificationUnknown
	}
	return c, nil
}
