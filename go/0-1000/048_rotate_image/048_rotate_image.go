package rotateimage

// from "rotation matrix" we know
// x = xCos(-Pi/2)-ySin(-Pi/2) = y
// y = xSin(-Pi/2)+yCos(-Pi/2) = -x

// If x is a row, and y is a column, then we need to shift matrix
// to the center of coordinates, rotate, shift back.
// xC = x - length / 2,
// yC = y - length / 2
// rotatedX = yC + length/2 = y
// rotatedY = -xC + length/2 = -x + length
func rotate(matrix [][]int) {
	size := len(matrix)

	for x := 0; x < size-1; x++ {
		for y := x; y < size-x-1; y++ {
			// sequentially move 4 cells to their new places
			value := matrix[x][y]
			for i := 0; i < 4; i++ {
				x2 := y
				y2 := -x + size - 1
				matrix[x2][y2], value = value, matrix[x2][y2]
				x, y = x2, y2
			}
		}
	}
}
