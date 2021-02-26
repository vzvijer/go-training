package anagram

import "strings"

func isAnagram(word string, candidate string) bool {
	word = strings.ToUpper(word)
	candidate = strings.ToUpper(candidate)
	w := []rune(word)
	c := []rune(candidate)

	if word == candidate || len(w) != len(c) {
		return false
	}

	letters := make(map[rune]int)
	for i := 0; i < len(w); i++ {
		letters[w[i]]++
		letters[c[i]]--
	}

	for _, n := range letters {
		if n != 0 {
			return false
		}
	}
	return true
}

func Detect(word string, candidates []string) []string {
	anagrams := []string{}
	for _, candidate := range candidates {
		if isAnagram(word, candidate) {
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}
