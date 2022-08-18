package reducearraysizetothehalf

import "sort"

func minSetSize(arr []int) int {
	countMap := map[int]int{}
	for _, n := range arr {
		countMap[n]++
	}

	uniqNums := make([]int, len(countMap))
	i := 0
	for k := range countMap {
		uniqNums[i] = k
		i++
	}

	sort.Slice(uniqNums, func(i, j int) bool {
		return countMap[uniqNums[i]] > countMap[uniqNums[j]]
	})

	countDeleted := 0
	for i, n := range uniqNums {
		countDeleted += countMap[n]
		if countDeleted >= len(arr)/2 {
			return i + 1
		}
	}

	return len(arr)
}
