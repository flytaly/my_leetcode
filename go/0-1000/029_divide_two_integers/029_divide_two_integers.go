package dividetoitegers

import "math"

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	sum := 0
	positive := true

	if divisor < 0 {
		divisor = -divisor
		positive = !positive
	}

	if dividend < 0 {
		dividend = -dividend
		positive = !positive
	}

	for s := 0; s <= dividend; s += divisor {
		sum += 1
	}

	if positive {
		return sum - 1
	}

	return 1 - sum
}
