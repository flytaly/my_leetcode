package containsduplicate

/*
	217. Contains Duplicate

	Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
*/

func containsDuplicate(nums []int) bool {
	visited := map[int]struct{}{}

	for _, n := range nums {
		if _, has := visited[n]; has {
			return true
		}
		visited[n] = struct{}{}
	}

	return false
}
