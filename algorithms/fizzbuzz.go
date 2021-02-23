package algorithms

import (
	"fmt"
	"strings"
)

func FizzBuzz(n int) string {
	var sb strings.Builder

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			sb.WriteString("Fizz Buzz")
		} else if i%3 == 0 {
			sb.WriteString("Fizz")
		} else if i%5 == 0 {
			sb.WriteString("Buzz")
		} else {
			sb.WriteString(fmt.Sprintf("%d", i))
		}

		if i != n {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("\n")

	return sb.String()
}
