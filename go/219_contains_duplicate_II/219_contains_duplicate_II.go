package contains_duplicate_ii

/*
	219. Contains Duplicate II

	Given an integer array nums and an integer k, return true if there are two distinct
	indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.
*/

func containsNearbyDuplicate(nums []int, k int) bool {
	indices := map[int][]int{}

	for i, num := range nums {
		ii := indices[num]
		for _, j := range ii {
			if i-j <= k { // i > j -> abs unnecessary
				return true
			}
		}
		indices[num] = append(indices[num], i)
	}

	return false
}
