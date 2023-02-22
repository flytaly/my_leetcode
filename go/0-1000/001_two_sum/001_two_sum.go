/*
https://leetcode.com/problems/two-sum/

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

*/

package twosums

func twoSum(nums []int, target int) []int {
	indices := map[int]int{}

	for i2, n2 := range nums {
		if i1, ok := indices[target-n2]; ok {
			return []int{i1, i2}
		}
		indices[n2] = i2
	}

	return []int{0, 0}
}
