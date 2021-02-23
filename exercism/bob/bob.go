// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	remark = strings.TrimSpace(remark)

	anyLetter := func() bool {
		for _, letter := range remark {
			if unicode.IsLetter(letter) {
				return true
			}
		}
		return false
	}()

	l := len(remark)
	if l == 0 {
		return "Fine. Be that way!"
	}

	if strings.ToUpper(remark) == remark && anyLetter {
		if remark[l-1] == '?' {
			return "Calm down, I know what I'm doing!"
		} else {
			return "Whoa, chill out!"
		}
	} else if remark[l-1] == '?' {
		return "Sure."
	}

	return "Whatever."
}
