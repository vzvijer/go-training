package algorithms

func Reverse(s string) string {
	rev := ""
	//this part of code doesn't work for non-English letters
	// for i := len(s) - 1; i >= 0; i-- {
	// 	rev = rev + string(s[i])
	// }
	for _, r := range []rune(s) {
		rev = string(r) + rev
	}
	return rev
}
