package reverse

import "fmt"

func Reverse(in string) string {
	letters := []rune(in)
	out := ""
	for _, letter := range letters {
		out = fmt.Sprintf("%c", letter) + out
	}
	return out
}
