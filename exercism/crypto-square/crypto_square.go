package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(text string) string {
	normsb := strings.Builder{}
	for _, letter := range []rune(text) {
		if unicode.IsLetter(letter) {
			normsb.WriteRune(unicode.ToLower(letter))
		}
	}
	dim := math.Round(math.Sqrt(normsb.Len()))
	nSpaces
	norms := normsb.String()

	
	cipher := strings.Builder{}
	for i := 0; i < dim; i++ {
		cipher.WriteString(norms[dim * i : dim * i + dim])
	}
}
