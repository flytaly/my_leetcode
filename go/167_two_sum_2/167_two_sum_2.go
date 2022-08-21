package two_sum_2

func twoSum(numbers []int, target int) []int {
	for i, n1 := range numbers {
		n2 := target - n1
		idx := binarySearch(numbers, n2, i+1)
		if idx == -1 {
			continue
		}
		return []int{i + 1, idx + 1}
	}

	return []int{}
}

func binarySearch(nums []int, target, startIdx int) int {
	left := startIdx
	right := len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid
			continue
		}
		left = mid + 1
	}

	if left < len(nums) && nums[left] == target {
		return left
	}

	return -1
}
