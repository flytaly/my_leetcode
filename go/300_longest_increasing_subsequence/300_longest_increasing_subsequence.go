/* 300. Longest Increasing Subsequence */

package longestincreasingsubsequence

// DP O(n^2) solution
func lengthOfLIS(nums []int) int {
	totalMax := 0
	lengths := make([]int, len(nums))

	for i, n := range nums {
		localMax := 1

		for j := 0; j < i; j++ {
			if nums[j] < n && localMax < lengths[j]+1 {
				localMax = lengths[j] + 1
			}
		}

		lengths[i] = localMax
		if totalMax < localMax {
			totalMax = localMax
		}
	}

	return totalMax
}
