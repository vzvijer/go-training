package pascal

func Triangle(n int) [][]int {
	if n == 0 {
		return [][]int{}
	} else {
		prevTriangle := Triangle(n - 1)
		row := make([]int, n)
		row[0] = 1

		if n > 1 {
			prevRow := prevTriangle[n-2]

			for i := 1; i < n-1; i++ {
				row[i] = prevRow[i-1] + prevRow[i]
			}
			row[n-1] = 1
		}

		return append(prevTriangle, row)
	}
}
