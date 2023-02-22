package house_robber

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func rob(nums []int) int {
	farSum, nearSum := 0, 0

	for _, n := range nums {
		next := farSum + n
		farSum = max(nearSum, farSum)
		nearSum = next
	}

	return max(farSum, nearSum)
}
