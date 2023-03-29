package minimumpathsum

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := m - 2; i >= 0; i-- { // fill the last column
		grid[i][n-1] += grid[i+1][n-1]
	}

	for j := n - 2; j >= 0; j-- { // fill the last row
		grid[m-1][j] += grid[m-1][j+1]
	}

	for i := m - 2; i >= 0; i-- { // from minimum mapth from the right-bottom corner
		for j := n - 2; j >= 0; j-- {
			grid[i][j] += min(grid[i][j+1], grid[i+1][j])
		}

	}
	return grid[0][0]
}
