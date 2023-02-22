package mergeintervals

import "sort"

/*
56. Merge Intervals

Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals,
and return an array of the non-overlapping intervals that cover all the intervals in the input.
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func mergeTwo(i1, i2 []int) []int {
	return []int{min(i1[0], i2[0]), max(i1[1], i2[1])}
}

func merge(intervals [][]int) [][]int {
	result := [][]int{}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return false
	})

	current := intervals[0]
	for _, next := range intervals[1:] {
		if current[1] >= next[0] {
			current = mergeTwo(current, next)
			continue
		}
		result = append(result, current)
		current = next
	}

	if len(current) != 0 {
		result = append(result, current)
	}

	return result
}
