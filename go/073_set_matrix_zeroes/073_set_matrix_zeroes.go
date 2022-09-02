package setmatrixzeroes

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	zeroCols := make([]bool, len(matrix[0]))

	nullifyRowLeft := func(row, col int) {
		for c := 0; c <= col; c++ {
			matrix[row][c] = 0
		}
	}
	nullifyColTop := func(row, col int) {
		for r := 0; r <= row; r++ {
			matrix[r][col] = 0
		}
	}

	for r, row := range matrix {
		nullifyRow := false
		for c, num := range row {
			if num == 0 {
				if !nullifyRow {
					nullifyRowLeft(r, c)
					nullifyRow = true
				}
				if !zeroCols[c] {
					nullifyColTop(r, c)
					zeroCols[c] = true
				}
				continue
			}
			if zeroCols[c] || nullifyRow {
				matrix[r][c] = 0
			}
		}
	}

}
