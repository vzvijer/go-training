package raindrops

import "fmt"

func Convert(n int) string {
	res := ""
	if n%3 == 0 {
		res = res + "Pling"
	}
	if n%5 == 0 {
		res = res + "Plang"
	}
	if n%7 == 0 {
		res = res + "Plong"
	}
	if res == "" {
		res = fmt.Sprintf("%d", n)
	}
	return res
}
