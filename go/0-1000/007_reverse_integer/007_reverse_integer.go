package reverseinteger

import (
	"math"
)

func reverse(x int) int {
	var result int32 = 0
	for x32 := int32(x); x32 != 0; x32 /= 10 {
		remainder := x32 % 10

		if x32 > 0 && result > (math.MaxInt32-remainder)/10 {
			return 0
		}
		if x32 <= 0 && result < (math.MinInt32-remainder)/10 {
			return 0
		}

		result = result*10 + remainder
	}

	return int(result)
}
