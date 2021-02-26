package sublist

type Relation string

func equal(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}
	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}

func isSublist(l1, l2 []int) bool {
	if len(l1) > len(l2) {
		return false
	}
	for i := 0; i <= len(l2)-len(l1); i++ {
		if equal(l1, l2[i:(i+len(l1))]) {
			return true
		}
	}
	return false
}

func Sublist(l1, l2 []int) Relation {
	len1 := len(l1)
	len2 := len(l2)
	if len1 == len2 && equal(l1, l2) {
		return "equal"
	} else if len1 < len2 && isSublist(l1, l2) {
		return "sublist"
	} else if len1 > len2 && isSublist(l2, l1) {
		return "superlist"
	} else {
		return "unequal"
	}
}
