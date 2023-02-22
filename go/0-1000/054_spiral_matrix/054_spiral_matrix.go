package spiralmatrix

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

func spiralOrder(matrix [][]int) []int {
	var result []int
	var col, row int

	moves := map[int](*Move){
		RIGHT:  &Move{&col, +1, len(matrix[0]) - 1},
		BOTTOM: &Move{&row, +1, len(matrix) - 1},
		LEFT:   &Move{&col, -1, 0},
		TOP:    &Move{&row, -1, 1},
	}

	dir := RIGHT
	end := len(matrix[0]) * len(matrix)
	for i := 0; i < end; i++ {
		result = append(result, matrix[row][col])
		if moves[dir].updateBound() {
			dir = (dir + 1) % 4
		}
		moves[dir].step()
	}

	return result
}
