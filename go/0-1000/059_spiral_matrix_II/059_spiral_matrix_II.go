package spiralmatrixII

type Move struct {
	pos   *int // reference to column or row index
	add   int  // direction +1 or -1
	bound int
}

func (m *Move) step() {
	*m.pos = *m.pos + m.add
}

func (m *Move) updateBound() bool {
	if *m.pos == m.bound {
		m.bound = *m.pos - m.add
		return true
	}
	return false
}

const (
	RIGHT = iota
	BOTTOM
	LEFT
	TOP
)

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	var col, row int
	dir := RIGHT

	moves := map[int](*Move){
		RIGHT:  &Move{&col, +1, n - 1},
		BOTTOM: &Move{&row, +1, n - 1},
		LEFT:   &Move{&col, -1, 0},
		TOP:    &Move{&row, -1, 1},
	}

	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	for i := 1; i <= n*n; i++ {
		if moves[dir].updateBound() {
			dir = (dir + 1) % 4
		}
		res[row][col] = i
		moves[dir].step()
	}

	return res
}
