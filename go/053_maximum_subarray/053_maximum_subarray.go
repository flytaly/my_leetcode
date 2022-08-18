package maximumsubarray

/* 53. Maximum Subarray */

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// divide and conquer
func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}

	left := maxSubArray(nums[:l/2])
	right := maxSubArray(nums[l/2:])

	center := nums[l/2-1] + nums[l/2]

	// <----
	for i := l/2 - 2; i >= 0; i-- {
		if center+nums[i] < center {
			break
		}
		center += nums[i]
	}
	// ---->
	for i := l/2 + 1; i < l; i++ {
		if center+nums[i] < center {
			break
		}
		center += nums[i]
	}

	maxSum := max(center, max(left, right))

	return maxSum
}

// iterations
// func maxSubArray(nums []int) int {
// 	maxSum := nums[0]
// 	curSum := nums[0]
//
// 	for _, num := range nums[1:] {
// 		if num > curSum+num {
// 			curSum = num
// 			maxSum = max(curSum, maxSum)
// 			continue
// 		}
// 		curSum += num
// 		maxSum = max(curSum, maxSum)
// 	}
//
// 	return maxSum
// }
