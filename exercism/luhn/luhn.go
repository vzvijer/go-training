package luhn

import (
	"strings"
	"unicode"
)

func Valid(cn string) bool {
	sum := 0
	n := 0
	digits := []rune(strings.TrimSpace(cn))
	for i := len(digits) - 1; i >= 0; i-- {
		d := digits[i]
		if unicode.IsDigit(d) {
			val := int(d - '0')
			n++
			if n%2 == 0 {
				val = (val * 2)
				if val > 9 {
					val = val - 9
				}
			}
			sum += val
		} else if unicode.IsSpace(d) {
			continue
		} else {
			return false
		}
	}
	if n <= 1 {
		return false
	}
	return sum%10 == 0
}
