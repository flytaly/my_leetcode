/* 300. Longest Increasing Subsequence */

package longestincreasingsubsequence

// DP with optimizations n*log(n)
func lengthOfLIS(nums []int) int {
	lastNums := make([]int, 1, len(nums)) // map length of a subsequence to it's last number
	lastNums[0] = nums[0]                 // index 0 means subsequence with length 1, 1=>2 ...

	for _, n := range nums {
		newLen := binarySearch(lastNums, n)

		if newLen > len(lastNums)-1 { // create new subsequence
			lastNums = append(lastNums, n)
			continue
		}
		// if 2 subsequences have the same length,
		// store only the one with the smallest last number
		if lastNums[newLen] > n {
			lastNums[newLen] = n
		}
	}

	return len(lastNums)
}

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid
			continue
		}
		left = mid + 1
	}

	return left
}

// DP O(n^2) solution
func lengthOfLIS_v1(nums []int) int {
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
