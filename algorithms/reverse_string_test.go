package algorithms

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	type testStruct struct {
		orig string
		rev  string
	}

	testData := []testStruct{
		{"", ""},
		{"a", "a"},
		{"abc", "cba"},
		{"XYZ", "ZYX"},
		{"abcDEF", "FEDcba"},
		{"Ћирибирибела", "алебирибириЋ"},
	}
	for i, test := range testData {
		t.Run(fmt.Sprintf("%d. test case", i+1), func(t *testing.T) {
			res := Reverse(test.orig)
			if res != test.rev {
				t.Fatalf("Reverse(\"%s\") got %s, wanted %s", test.orig, res, test.rev)
			}
		})
	}
}
