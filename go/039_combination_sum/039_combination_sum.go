package combinationsum

import "sort"

func combSum(candidates []int, target int, start int, result *([][]int), stack []int) {
	if target < 0 {
		return
	}
	for i := start; i < len(candidates); i++ {
		num := candidates[i]
		if num == target {
			s := make([]int, len(stack)+1)
			copy(s, stack)
			s[len(stack)] = num
			*result = append(*result, s)
			continue
		}
		stack = append(stack, num)
		combSum(candidates, target-num, i, result, stack)
		stack = stack[:len(stack)-1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	sort.Slice(candidates, func(i, j int) bool { return candidates[i] < candidates[j] })

	combSum(candidates, target, 0, &result, []int{})

	return result
}

func combinationSumV1(candidates []int, target int) [][]int {

	result := [][]int{}

	for i, num := range candidates {
		if num > target {
			continue
		}

		if num == target {
			result = append(result, []int{num})
			continue
		}

		sums := combinationSumV1(candidates[i:], target-num)

		if len(sums) == 0 {
			continue
		}

		for _, sum := range sums {
			s := []int{num}
			for _, n := range sum {
				s = append(s, n)
			}
			result = append(result, s)
		}
	}

	return result
}
