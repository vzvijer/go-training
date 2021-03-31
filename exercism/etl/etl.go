package etl

import "strings"

func Transform(input map[int][]string) map[string]int {
	output := map[string]int{}
	for ivalue, iletters := range input {
		for _, oletter := range iletters {
			output[strings.ToLower(oletter)] = ivalue
		}
	}
	return output
}
