package jumpgame

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func canJump(nums []int) bool {
	maxIndex := 0

	for i, v := range nums {
		if i > maxIndex {
			return false
		}
		maxIndex = max(maxIndex, v+i)
	}

	return true
}
