package searchinrotatedsortedarray

/*
	33. Search in Rotated Sorted Array
	https://leetcode.com/problems/search-in-rotated-sorted-array/
*/

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l < r {
		m := (l + r) / 2

		if target == nums[m] {
			return m
		}

		// left segment is sorted
		if nums[l] <= nums[m] {
			if target >= nums[l] && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
			continue
		}

		// otherwise right segment should be sorted
		if target > nums[m] && target <= nums[r] {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	if nums[l] == target {
		return l
	}

	return -1
}
