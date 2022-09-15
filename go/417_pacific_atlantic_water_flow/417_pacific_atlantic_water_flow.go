package pacificatlanticwaterflow

type State byte

const (
	NotVisited State = iota
	Reachable
	Unreachable
	Visiting
)

func dfs(i, j, dir int, heights *[][]int, ocean *[][]State) State {

	if (*ocean)[i][j] == Reachable || (*ocean)[i][j] == Unreachable {
		return (*ocean)[i][j]
	}

	// all 4 directions, clockwise if dir==+1, or counterclockwise if dir==-1
	dirs := [][]int{{i + dir, j}, {i, j + dir}, {i - dir, j}, {i, j - dir}}

	for _, d := range dirs {
		x, y := d[0], d[1]
		// outside of the matrix
		if x >= len(*heights) || x < 0 || y >= len((*heights)[0]) || y < 0 {
			continue
		}
		// skip currently visiting cell or if height value of the cell is bigger
		if (*ocean)[x][y] == Visiting || (*heights)[i][j] < (*heights)[x][y] {
			continue
		}

		if (*ocean)[x][y] == Reachable {
			return Reachable
		}

		if (*ocean)[x][y] == NotVisited {
			(*ocean)[x][y] = Visiting
			(*ocean)[x][y] = dfs(x, y, dir, heights, ocean)
			if (*ocean)[x][y] == Reachable {
				return Reachable
			}
		}

	}

	return Unreachable
}

func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	pacific := make([][]State, m)  // can reach Pacific Ocean
	atlantic := make([][]State, m) // can reach Atlantic Ocean

	// Fill left and right borders
	for i := 0; i < m; i++ {
		pacific[i] = make([]State, n)
		pacific[i][0] = Reachable
		atlantic[m-1-i] = make([]State, n)
		atlantic[m-1-i][n-1] = Reachable
	}
	// Fill top and bottom borders
	for j := 1; j < len(heights[0]); j++ {
		pacific[0][j] = Reachable
		atlantic[m-1][n-1-j] = Reachable
	}

	// Fill everyting else
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			pacific[i][j] = dfs(i, j, -1, &heights, &pacific)
			atlantic[m-1-i][n-1-j] = dfs(m-1-i, n-1-j, +1, &heights, &atlantic)
		}
	}

	result := [][]int{}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacific[i][j] == Reachable && atlantic[i][j] == Reachable {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}
