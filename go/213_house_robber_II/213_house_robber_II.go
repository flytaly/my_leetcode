package house_robber_II

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func robV1(nums []int) int {
	farSum, nearSum := 0, 0

	for _, n := range nums {
		next := farSum + n
		farSum = max(nearSum, farSum)
		nearSum = next
	}

	return max(farSum, nearSum)
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(
		robV1(nums[1:]),
		robV1(nums[:len(nums)-1]),
	)
}
