package algorithms

import (
	"fmt"
	"testing"
)

func TestDecToBase(t *testing.T) {
	type testStruct struct {
		n, b int
		out  string
	}

	testData := []testStruct{
		{10, 2, "1010"},
		{255, 16, "FF"},
	}

	for i, test := range testData {
		t.Run(fmt.Sprintf("%d. test case", i+1), func(t *testing.T) {
			res, err := DecToBase(test.n, test.b)
			if err != nil {
				t.Errorf(err.Error())
			}
			if res != test.out {
				t.Fatalf("%d(10) -> %s(%d), expected %s(%d)", test.n, res, test.b, test.out, test.b)
			}
		})
	}
}
