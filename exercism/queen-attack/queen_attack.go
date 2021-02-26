package queenattack

import (
	"errors"
)

func absDistance(r1, r2 rune) rune {
	distance := r1 - r2
	if distance < 0 {
		distance = -distance
	}
	return distance
}

func CanQueenAttack(q1, q2 string) (bool, error) {
	if len(q1) > 2 || len(q2) > 2 || q1 == q2 {
		return false, errors.New("Wrong queen position.")
	}

	c1, r1 := []rune(q1)[0], []rune(q1)[1]
	c2, r2 := []rune(q2)[0], []rune(q2)[1]

	if (c1 < 'a' || c1 > 'h') || (r1 < '1' || r1 > '8') ||
		(c2 < 'a' || c2 > 'h') || (r2 < '1' || r2 > '8') {
		return false, errors.New("Wrong queen position.")
	}

	return c1 == c2 || r1 == r2 || absDistance(c1, c2) == absDistance(r1, r2), nil
}
