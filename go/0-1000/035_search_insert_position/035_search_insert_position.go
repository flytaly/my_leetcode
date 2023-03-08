package searchinsert

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid - 1
			continue
		}
		l = mid + 1
	}

	return l
}
