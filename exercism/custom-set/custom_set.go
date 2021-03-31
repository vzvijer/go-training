package stringset

import "fmt"

type Set map[string]bool

func (set Set) String() string {
	s := "{"
	delimiter := ""
	for k := range set {
		s = s + delimiter + fmt.Sprintf("\"%s\"", k)
		delimiter = ", "
	}
	s = s + "}"
	return s
}

func (set Set) IsEmpty() bool {
	return len(set) == 0
}

func (set Set) Has(item string) bool {
	return set[item]
}

func (set Set) Add(item string) {
	(set)[item] = true
}

func New() Set {
	return map[string]bool{}
}

func NewFromSlice(slice []string) Set {
	set := New()
	for _, item := range slice {
		set[item] = true
	}
	return set
}

func Subset(subset, set Set) bool {
	for k := range subset {
		if !set.Has(k) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	return Subset(s1, s2)
}

func Intersection(s1, s2 Set) Set {
	intersection := Set{}
	for k := range s1 {
		if s2.Has(k) {
			intersection[k] = true
		}
	}
	return intersection
}

func Difference(s1, s2 Set) Set {
	difference := Set{}
	for k := range s1 {
		if !s2.Has(k) {
			difference[k] = true
		}
	}
	return difference
}

func Union(s1, s2 Set) Set {
	union := Set{}
	for k := range s1 {
		union[k] = true
	}
	for k := range s2 {
		union[k] = true
	}
	return union
}

func Disjoint(s1, s2 Set) bool {
	for k := range s1 {
		if s2[k] == true {
			return false
		}
	}
	return true
}
