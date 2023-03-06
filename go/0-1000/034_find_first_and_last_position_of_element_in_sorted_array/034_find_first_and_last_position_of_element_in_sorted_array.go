package findfirstandlastpos

// Find the first position of a target in an array.
// If the array doesn't include target, then the function
// returns the position where the target supposed to be.
func findFirstPos(nums []int, target int) int {
	l, r := 0, len(nums)-1
	first := len(nums)

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			first = mid
			r = mid - 1
			continue
		}
		l = mid + 1
	}

	return first
}

func searchRange(nums []int, target int) []int {
	first := findFirstPos(nums, target)

	if first == len(nums) || nums[first] != target {
		return []int{-1, -1}
	}

	last := findFirstPos(nums, target+1) - 1

	return []int{first, last}
}
