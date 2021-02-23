package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(ss []string) FreqMap {
	res := FreqMap{}

	c := make(chan FreqMap)
	defer close(c)

	for i := 0; i < len(ss); i++ {
		go func(s string) {
			c <- Frequency(s)
		}(ss[i])
	}

	for i := 0; i < len(ss); i++ {
		m := <-c
		for k, v := range m {
			res[k] += v
		}
	}

	return res
}
