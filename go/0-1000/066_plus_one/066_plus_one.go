package plusone

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = (digits[i] + 1) % 10
		if digits[i] != 0 {
			break
		}
	}

	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}

	// if digits[0] == 0 {
	// 	result := make([]int, len(digits)+1)
	// 	result[0] = 1
	// 	copy(result[1:], digits)
	// 	return result
	// }

	return digits
}
