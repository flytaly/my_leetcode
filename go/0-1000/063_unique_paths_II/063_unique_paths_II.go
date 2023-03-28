package uniquepathsII

/*
   An example of the solution table for 4x2 case without obstacles.
   +----+----+----+----+
   | 10 | 06 | 03 | 01 |
   | 04 | 03 | 02 | 01 |
   | 01 | 01 | 01 | -- |
   +----+----+----+----+
*/

func uniquePathsWithObstacles(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] != 0 {
				grid[i][j] = 0 // reuse the same grid and mark obstacles as 0
				continue
			}

			if i == m-1 && j == n-1 { // bottom-right corner
				grid[i][j] = 1
				continue
			}

			right, bottom := 0, 0 // right and bottom neighbours
			if i < m-1 {
				right = grid[i+1][j]
			}
			if j < n-1 {
				bottom = grid[i][j+1]
			}
			grid[i][j] = right + bottom
		}
	}

	return grid[0][0]
}
