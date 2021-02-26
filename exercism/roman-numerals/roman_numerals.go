package romannumerals

import (
	"errors"
	"strings"
)

func ToRomanNumeral(arabic int) (string, error) {
	arabics := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := strings.Builder{}
	if arabic < 1 || arabic > 3000 {
		return "", errors.New("Only values between 1 and 3000 can be represented by romans")
	}
	for arabic > 0 {
		for i := 0; i < len(arabics); i++ {
			if arabic >= arabics[i] {
				roman.WriteString(romans[i])
				arabic = arabic - arabics[i]
				i--
			}
		}
	}
	return roman.String(), nil
}
