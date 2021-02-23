// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"fmt"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	res := ""

	letters := []rune(s)
	prevLetter := ' '
	for _, letter := range letters {
		if unicode.IsLetter(letter) && !unicode.IsLetter(prevLetter) {
			if prevLetter == '\'' && unicode.ToUpper(letter) == 'S' {
			} else {
				res = res + fmt.Sprintf("%c", unicode.ToUpper(letter))
			}
		}
		prevLetter = letter
	}

	return res
}
