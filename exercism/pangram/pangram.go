package pangram

import "unicode"

func IsPangram(text string) bool {
	usedLetters := map[rune]bool{}
	for _, letter := range []rune(text) {
		if unicode.IsLetter(letter) {
			usedLetters[unicode.ToLower(letter)] = true
		}
	}
	return len(usedLetters) == 26
}
