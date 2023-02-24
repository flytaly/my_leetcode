package zigzag_conversion

import "strings"

/*
3 rows
P   A   H   N
A P L S I I G
Y   I   R
1 -> +4
2 -> +2 +2
3 -> +4
============
4 rows

P     I    N
A   L S  I G
Y A   H R
P     I

1 -> +6
2 -> +4 +2
3 -> +2 +4
4 -> +6
*/

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var sb strings.Builder
	loopLen := numRows + numRows - 2

	for row := 0; row < numRows; row++ {
		i := row
		direction := 1 // alternate between 1 and 0 and represet moving direction
		for {
			if i >= len(s) {
				break
			}
			sb.WriteByte(s[i])
			next := direction*(numRows-row-1) + (1-direction)*(row)
			direction = 1 - direction
			if next != 0 {
				i += 2 * next
				continue
			}
			if loopLen == 0 {
				break
			}
			i += loopLen
		}
	}

	return sb.String()
}
