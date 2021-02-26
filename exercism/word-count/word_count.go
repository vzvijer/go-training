package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(s string) Frequency {
	freq := Frequency{}
	r, _ := regexp.Compile("\\w+(['_-]?\\w+)?")
	words := r.FindAllString(s, -1)
	for _, word := range words {
		freq[strings.ToLower(word)]++
	}
	return freq
}
