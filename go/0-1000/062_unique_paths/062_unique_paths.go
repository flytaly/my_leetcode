package uniquepaths

/*
   62. Unique Paths

   An example of the solution table for 4x2 case
   +----+----+----+----+
   | 10 | 06 | 03 | 01 |
   | 04 | 03 | 02 | 01 |
   | 01 | 01 | 01 | -- |
   +----+----+----+----+
*/

func uniquePaths(m int, n int) int {
	column := make([]int, m) // moving column from right to left
	column[0] = 1            // count indices from bottom

	for colIndex := 0; colIndex < n; colIndex++ {
		for i := range column[1:] {
			column[i+1] += column[i]
		}
	}

	return column[len(column)-1]
}
