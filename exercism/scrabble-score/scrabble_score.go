package scrabble

import "unicode"

func Score(word string) int {
	letterValues := []int{1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10}
	score := 0
	for _, letter := range word {
		firstLetter := 'A'
		if unicode.IsLower(letter) {
			firstLetter = 'a'
		}
		score = score + letterValues[letter-firstLetter]
	}
	return score
}
