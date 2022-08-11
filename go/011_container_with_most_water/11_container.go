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

type Line struct {
	h int
	i int
}

func maxArea(height []int) int {
	lines := []Line{{height[0], 0}}
	result := 0

	for i := 1; i < len(height); i++ {
		h := height[i]

		for _, l := range lines {
			minH := l.h
			if minH > h {
				minH = h
			}
			a := minH * (i - l.i)

			if a > result {
				result = a
			}

		}

		length := len(lines)
		if lines[length-1].h < h {
			lines = append(lines, Line{h, i})
		}

	}
	// fmt.Println(lines)
	return result
}
