package insertinterval

/*
57. Insert Interval
You are given an array of non-overlapping intervals intervals where intervals[i] = [starti, endi] represent the start and the end of the ith interval and intervals is sorted in ascending order by starti. You are also given an interval newInterval = [start, end] that represents the start and end of another interval.

Insert newInterval into intervals such that intervals is still sorted in ascending order by starti and intervals still does not have any overlapping intervals (merge overlapping intervals if necessary).

Return intervals after the insertion.
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

func merge(i1, i2 []int) []int {
	return []int{min(i1[0], i2[0]), max(i1[1], i2[1])}
}

func insert(intervals [][]int, newInterval []int) [][]int {
	output := [][]int{}

	inserted := false
	for _, inter := range intervals {
		if inter[1] < newInterval[0] { // intervals don't intersect
			output = append(output, inter)
			continue
		}

		if inserted {
			output = append(output, inter)
			continue
		}

		// intervals intersect => merge
		if inter[0] <= newInterval[1] {
			newInterval = merge(inter, newInterval)
			continue
		}

		output = append(output, newInterval, inter) // insert new interval
		inserted = true
	}

	if !inserted {
		output = append(output, newInterval) // insert new interval
	}

	return output
}
