package algorithms

import (
	"fmt"
	"testing"
)

type testStruct struct {
	list []int
	sum  int
}

func TestSum(t *testing.T) {
	testData := []testStruct{
		{nil, 0},
		{[]int{}, 0},
		{[]int{1}, 1},
	}

	for _, test := range testData {
		t.Run(fmt.Sprintf("%v -> %v", test.list, test.sum), func(t *testing.T) {
			res := Sum(test.list)
			if test.sum != res {
				t.Fatalf("Sum(%v) returned %v, wanted %v", test.list, res, test.sum)
			}
		})
	}
}
