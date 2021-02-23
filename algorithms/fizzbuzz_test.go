package algorithms

import (
	"fmt"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	type testStruct struct {
		n   int
		out string
	}

	testData := []testStruct{
		{1, "1\n"},
		{2, "1, 2\n"},
		{3, "1, 2, Fizz\n"},
		{4, "1, 2, Fizz, 4\n"},
		{5, "1, 2, Fizz, 4, Buzz\n"},
		{6, "1, 2, Fizz, 4, Buzz, Fizz\n"},
		{7, "1, 2, Fizz, 4, Buzz, Fizz, 7\n"},
		{8, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8\n"},
		{9, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz\n"},
		{10, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz\n"},
		{11, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11\n"},
		{12, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz\n"},
		{13, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13\n"},
		{14, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14\n"},
		{15, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz\n"},
		{16, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz, 16\n"},
		{17, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz, 16, 17\n"},
		{18, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz, 16, 17, Fizz\n"},
		{19, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz, 16, 17, Fizz, 19\n"},
		{20, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz, 16, 17, Fizz, 19, Buzz\n"},
	}

	for i, test := range testData {
		t.Run(fmt.Sprintf("%d. test case", i+1), func(t *testing.T) {
			res := FizzBuzz(test.n)
			if res != test.out {
				t.Fatalf("%d: expected %s, got %s", test.n, test.out, res)
			}
		})
	}
}
