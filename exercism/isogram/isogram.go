package isogram

import "unicode"

func IsIsogram(s string) bool {
	var usedLetters [26]bool
	letters := []rune(s)
	for _, letter := range letters {
		if unicode.IsLetter(letter) {
			letter = unicode.ToLower(letter)
			if usedLetters[letter-'a'] {
				return false
			}
			usedLetters[letter-'a'] = true
		}
	}
	return true
}
