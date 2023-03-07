package validsudoku

func isValidSudoku(board [][]byte) bool {
	var digits map[byte]bool

	var clear = func() {
		digits = make(map[byte]bool, 9)
	}

	var isValid = func(d byte) bool {
		if d == '.' {
			return true
		}
		if d < '1' || d > '9' || digits[d] {
			return false
		}
		digits[d] = true
		return true
	}

	// check rows
	for _, row := range board {
		clear()
		for _, d := range row {
			if !isValid(d) {
				return false
			}
		}
	}

	// check columns
	for i := 0; i < 9; i++ {
		clear()
		for _, row := range board {
			if !isValid(row[i]) {
				return false
			}
		}

	}
	// check 3x3 boxes
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			clear()
			for ii := i; ii < i+3; ii++ {
				for jj := j; jj < j+3; jj++ {
					if !isValid(board[ii][jj]) {
						return false
					}
				}
			}
		}
	}

	return true
}
