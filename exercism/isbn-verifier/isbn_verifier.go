package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	sum := 0
	nums := []rune(strings.ReplaceAll(isbn, "-", ""))

	if len(nums) != 10 {
		return false
	}

	for i := 0; i < 10; i++ {
		if i == 9 && nums[i] == 'X' {
			sum += 10
		} else {
			if n, err := strconv.Atoi(string(nums[i])); err != nil {
				return false
			} else {
				sum += (10 - i) * n
			}
		}
	}
	return sum%11 == 0
}
