package findminimuminrotatedsortedarray

/* 153. Find Minimum in Rotated Sorted Array */

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	minVal := nums[0]

	for left < right {
		middle := (left + right) / 2
		if nums[left] <= nums[middle] {
			left = middle + 1
			minVal = min(minVal, nums[left])
		} else {
			right = middle - 1
			minVal = min(minVal, nums[middle])
		}
	}

	return minVal
}
