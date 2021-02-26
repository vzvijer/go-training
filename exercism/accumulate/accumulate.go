package accumulate

func Accumulate(words []string, f func(string) string) []string {
	res := make([]string, len(words))
	for i, word := range words {
		res[i] = f(word)
	}
	return res
}
