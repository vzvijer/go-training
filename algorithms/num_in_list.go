package algorithms

func NumInList(l []int, n int) bool {
	if l != nil {
		for _, v := range l {
			if v == n {
				return true
			}
		}
	}
	return false
}
