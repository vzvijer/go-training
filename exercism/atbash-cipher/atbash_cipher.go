package atbash

import (
	"strings"
	"unicode"
)

func Atbash(text string) string {
	cipher := strings.Builder{}
	i := 0
	for _, r := range []rune(text) {
		if i == 5 {
			cipher.WriteRune(' ')
			i = 0
		}
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			continue
		} else {
			if unicode.IsUpper(r) {
				cipher.WriteRune('z' - unicode.ToLower(r) + 'a')
			} else if unicode.IsLower(r) {
				cipher.WriteRune('z' - r + 'a')
			} else {
				cipher.WriteRune(r)
			}
			i++
		}
	}
	return strings.TrimSpace(cipher.String())
}
