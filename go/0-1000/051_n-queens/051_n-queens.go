package nqueens

import "strings"

type Board struct {
	cols     map[int]bool
	negative map[int]bool // negative diagonals
	positive map[int]bool // positive diagonals
}

func newBoard() *Board {
	return &Board{cols: map[int]bool{}, positive: map[int]bool{}, negative: map[int]bool{}}
}

func (b *Board) isValidPlace(row, col int) bool {
	return !b.cols[col] && !b.negative[col-row] && !b.positive[col+row]
}

func (b *Board) set(row, col int) {
	b.cols[col], b.negative[col-row], b.positive[col+row] = true, true, true
}

func (b *Board) clear(row, col int) {
	b.cols[col], b.negative[col-row], b.positive[col+row] = false, false, false
}

func solveNQueens(n int) [][]string {
	var result [][]string
	var board = newBoard()
	var dfs func(prev *[]string)

	dfs = func(rows *[]string) {
		row := len(*rows)
		for col := 0; col < n; col++ {
			if !board.isValidPlace(row, col) {
				continue
			}
			str := strings.Repeat(".", col) + "Q" + strings.Repeat(".", n-col-1)
			*rows = append(*rows, str)
			board.set(row, col)

			if row < n-1 {
				dfs(rows)
			}
			if row == n-1 {
				res := make([]string, len(*rows))
				copy(res, *rows)
				result = append(result, res)
			}
			board.clear(row, col)
			*rows = (*rows)[:len(*rows)-1]
		}
	}

	dfs(&[]string{})

	return result
}
