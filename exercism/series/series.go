package series

func UnsafeFirst(n int, s string) string {
	if n > len(s) {
		return s
	}

	return string([]rune(s)[:n])
}

func All(n int, s string) []string {
	if n > len(s) {
		return nil
	}

	k := len(s) + 1 - n
	all := make([]string, k)
	for i := 0; i < k; i++ {
		all[i] = UnsafeFirst(n, string([]rune(s)[i:i+n]))
	}
	return all
}
