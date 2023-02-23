package medianoftwosortedarrays

import (
	"math"
)

/*
	A = {1,2,3,4}
    B = {1,2,3,4,5,6,7,8}
	Total lenght is 4+8=12, so half = 6.
    split A to 2 even parts => 1 2 | 3 4
    split B as 6 - 2 = 4 => 1 2 3 4 | 5 6 7 8

	Check if left part of A is also left part of B, if true then we found median,
	otherwise increase or decrease partion and repeat the process.
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	total := len(nums1) + len(nums2)
	half := total / 2

	l, r := 0, len(nums1)-1

	for {
		var midA int
		if l > r {
			midA = -1
		} else {
			midA = (l + r) / 2
		}
		midB := half - midA - 2

		var ALeft, ARight, BLeft, BRight float64

		if midA < 0 || len(nums1) == 0 {
			ALeft = math.Inf(-1)
		} else {
			ALeft = float64(nums1[midA])
		}

		if midA+1 >= len(nums1) {
			ARight = math.Inf(1)
		} else {
			ARight = float64(nums1[midA+1])
		}

		if midB < 0 {
			BLeft = math.Inf(-1)
		} else {
			BLeft = float64(nums2[midB])
		}

		if midB+1 >= len(nums2) {
			BRight = math.Inf(1)
		} else {
			BRight = float64(nums2[midB+1])
		}

		if ALeft <= BRight && BLeft <= ARight {
			if total%2 != 0 {
				return math.Min(ARight, BRight)
			}
			return (math.Max(ALeft, BLeft) + math.Min(ARight, BRight)) / 2
		}

		if ALeft > BRight {
			r = midA - 1
			continue
		}
		l = midA + 1
	}
}
