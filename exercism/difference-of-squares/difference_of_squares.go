package diffsquares

func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

func Difference(n int) int {
	res := SumOfSquares(n) - SquareOfSum(n)
	if res < 0 {
		res = -res
	}
	return res
}
