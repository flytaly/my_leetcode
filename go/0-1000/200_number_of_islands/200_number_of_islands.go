package numberofislands

const (
	Water   = '0'
	Land    = '1'
	Visited = '2'
)

func dfs(i, j int, grid *[][]byte) {
	if i < 0 || j < 0 || i >= len(*grid) || j >= len((*grid)[0]) {
		return
	}

	if (*grid)[i][j] != Land {
		return
	}

	(*grid)[i][j] = Visited

	dfs(i+1, j, grid)
	dfs(i, j+1, grid)
	dfs(i-1, j, grid)
	dfs(i, j-1, grid)
}

func numIslands(grid [][]byte) (result int) {

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == Land {
				dfs(i, j, &grid)
				result += 1
			}
		}
	}

	return result
}
