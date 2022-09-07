package missingnumber

import "sort"

// v4:  XOR
// Using XOR multiple times we can eliminate duplicates because n^n=0 and a^b^a=b.
func missingNumber(nums []int) int {
	result := len(nums)

	for i := 0; i < len(nums); i++ {
		result ^= i ^ nums[i]
	}

	return result
}

// v3: arithmetic series => n(a1+an)/2
func missingNumberSum(nums []int) int {
	sumtotal := len(nums) * (len(nums) + 1) / 2
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sumtotal - sum
}

// v2: hash map
func missingNumberHash(nums []int) int {
	hash := map[int]struct{}{}

	for _, v := range nums {
		hash[v] = struct{}{}
	}

	for i := 0; i < len(nums); i++ {
		if _, exist := hash[i]; !exist {
			return i
		}
	}

	return len(nums)
}

// v1: sort O(n*logn+n)
func missingNumberSort(nums []int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	for i, v := range nums {
		if v != i {
			return i
		}
	}
	return len(nums)
}
