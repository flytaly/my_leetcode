package nonoverlapping_intervals

import (
	"sort"
)

/*
435. Non-overlapping Intervals

Given an array of intervals intervals where intervals[i] = [starti, endi], return the minimum number of intervals
you need to remove to make the rest of the intervals non-overlapping.
*/

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return false
	})

	deleteCount := 0
	current := intervals[0]
	for i := 1; i < len(intervals); i++ {
		next := intervals[i]
		// no intersect => go to the next step
		if current[1] <= next[0] {
			current = next
			continue
		}
		// intersect => remove an element with bigger right boundary
		if current[1] > next[1] {
			current = next
		}
		deleteCount++
	}

	return deleteCount
}
