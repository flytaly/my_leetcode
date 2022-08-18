package maximumsubarray

/* 53. Maximum Subarray */

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	curSum := nums[0]

	for _, num := range nums[1:] {
		if num > curSum+num {
			curSum = num
			maxSum = max(curSum, maxSum)
			continue
		}
		curSum += num
		maxSum = max(curSum, maxSum)
	}

	return maxSum
}
