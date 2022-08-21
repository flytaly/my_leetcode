package threeSum

import "sort"

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
		usedNums := map[int]int{}
		for ii := len(nums) - 1; ii > i; ii-- { // loop over numbers from the end
			last := nums[ii]
			if last < 0 {
				break
			}
			if usedNums[last] > 0 || numMap[last] <= 0 {
				continue
			}
			usedNums[last]++

			diff := -last - first // the middle number
			if numMap[diff]-usedNums[diff] <= 0 {
				continue
			}
			usedNums[diff]++
			result = append(result, []int{first, diff, last})
		}
		numMap[first] = 0
	}

	return result
}
