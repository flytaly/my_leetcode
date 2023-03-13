package nqueensii

func totalNQueens(n int) (result int) {
	cols, negative, positive := map[int]bool{}, map[int]bool{}, map[int]bool{}

	var dfs func(row int)

	dfs = func(row int) {
		for col := 0; col < n; col++ {
			if cols[col] || negative[col-row] || positive[col+row] {
				continue
			}
			cols[col], negative[col-row], positive[col+row] = true, true, true

			if row < n-1 {
				dfs(row + 1)
			}

			if row == n-1 {
				result += 1
			}

			cols[col], negative[col-row], positive[col+row] = false, false, false
		}
	}

	dfs(0)

	return result
}
