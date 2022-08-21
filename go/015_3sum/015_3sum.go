package threeSum

import (
	"math"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	result := [][]int{}

	numMap := map[int]int{} // store how many times a number appears
	for _, n := range nums {
		numMap[n]++
	}

	for i, first := range nums {
		if first > 0 { // first number should be not positive
			return result
		}
		if numMap[first] == 0 {
			continue
		}
		numMap[first] -= 1
		prev := math.MinInt
		for ii := len(nums) - 1; ii > i; ii-- { // loop over numbers from the end
			last := nums[ii]
			if last < 0 {
				break
			}
			if numMap[last] < 0 {
				continue
			}
			diff := -last - first // the middle number
			if diff == prev {
				continue
			}
			if diff > last {
				break
			}
			numMap[last]--
			if numMap[diff] > 0 {
				prev = diff
				result = append(result, []int{first, diff, last})
			}
			numMap[last]++
		}
		numMap[first] = 0
	}

	return result
}
