package threeSum

import (
	"math"
	"sort"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func threeSumClosest(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var sum = math.MaxInt

	for i := 0; i < len(nums); i++ {
		pL, pR := i+1, len(nums)-1 // left (next to the current number) and right pointers
		for pL < pR {
			nextSum := nums[i] + nums[pL] + nums[pR]
			diff := target - nextSum
			if abs(diff) < abs(target-sum) {
				sum = nextSum
			}
			if diff == 0 {
				return target
			}
			if diff > 0 {
				pL += 1
			}
			if diff < 0 {
				pR -= 1
			}
		}
	}

	return sum
}
