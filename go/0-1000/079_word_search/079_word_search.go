package wordsearch

var DIRS = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func search(i, j int, board *[][]byte, word string, visited *[][]bool) bool {
	if len(word) == 0 {
		return true
	}

	if (*board)[i][j] != word[0] || (*visited)[i][j] {
		return false
	}

	if len(word) == 1 {
		return true
	}

	(*visited)[i][j] = true
	defer func() { (*visited)[i][j] = false }()

	for _, dir := range DIRS {
		ii, jj := i+dir[0], j+dir[1]
		if ii >= len(*board) || ii < 0 || jj >= len((*board)[0]) || jj < 0 {
			continue
		}
		if search(ii, jj, board, word[1:], visited) {
			return true
		}
	}

	return false
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	for i := range board {
		for j := range board[i] {
			if search(i, j, &board, word, &visited) {
				return true
			}
		}
	}

	return false
}
