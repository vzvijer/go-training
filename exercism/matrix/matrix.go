package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(in string) (Matrix, error) {
	var n_rows, n_cols int

	m := Matrix{}
	rows := strings.Split(in, "\n")
	n_rows = len(rows)
	m = make([][]int, n_rows)
	for i := 0; i < n_rows; i++ {
		row := strings.TrimSpace(rows[i])
		cells := strings.Split(row, " ")
		if i == 0 {
			n_cols = len(cells)
		} else {
			if n_cols != len(cells) {
				return nil, errors.New("Wrong columns number")
			}
		}
		m[i] = make([]int, n_cols)
		for j := 0; j < n_cols; j++ {
			cell := cells[j]
			val, err := strconv.Atoi(cell)
			if err != nil {
				return nil, errors.New(err.Error())
			}
			m[i][j] = val
		}
	}
	return m, nil
}

func (m Matrix) Set(r, c, val int) bool {
	var n_rows, n_cols int

	n_rows = len(m)
	if n_rows > 0 {
		n_cols = len(m[0])
	}
	if r < 0 || c < 0 || r >= n_rows || c >= n_cols {
		return false
	}
	m[r][c] = val
	return true
}

func (m Matrix) Rows() [][]int {
	var n_rows, n_cols int

	n_rows = len(m)
	if n_rows > 0 {
		n_cols = len(m[0])
	}

	res := make([][]int, n_rows)
	for i := 0; i < n_rows; i++ {
		res[i] = make([]int, n_cols)
		for j := 0; j < n_cols; j++ {
			res[i][j] = m[i][j]
		}
	}
	return res
}

func (m Matrix) Cols() [][]int {
	var n_rows, n_cols int

	n_rows = len(m)
	if n_rows > 0 {
		n_cols = len(m[0])
	}

	res := make([][]int, n_cols)
	for i := 0; i < n_cols; i++ {
		res[i] = make([]int, n_rows)
		for j := 0; j < n_rows; j++ {
			res[i][j] = m[j][i]
		}
	}
	return res
}
