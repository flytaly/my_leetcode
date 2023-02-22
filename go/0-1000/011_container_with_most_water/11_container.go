/*
11. Container With Most Water

https://leetcode.com/problems/container-with-most-water/

You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the i-th line are (i, 0) and (i, height[i]).
Find two lines that together with the x-axis form a container, such that the container contains the most water.
Return the maximum amount of water a container can store.
Notice that you may not slant the container.

Example 1:

Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49

Example 2:

Input: height = [1,1]
Output: 1

Constraints:

n == height.length
2 <= n <= 105
0 <= height[i] <= 104

*/

package container_with_most_water

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxArea(height []int) int {
	result := 0
	left := 0
	right := len(height) - 1
	length := right - left
	for ; length > 0; length = right - left {
		width := min(height[left], height[right])
		area := width * length
		if area > result {
			result = area
		}
		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}
	}

	return result
}
