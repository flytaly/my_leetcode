package coinchange

import "sort"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChange(coins []int, amount int) int {

	sort.Slice(coins, func(i, j int) bool { return coins[i] < coins[j] })

	amounts := make([]int, amount+1)
	amounts[0] = 0
	maxVal := amount + 1
	for i := 1; i <= amount; i++ { // fill with max value
		amounts[i] = maxVal
	}

	for i := coins[0]; i <= amount; i++ {
		for _, v := range coins {
			if v == i {
				amounts[i] = 1
				break
			}
			if v > i {
				break
			}
			amounts[i] = min(1+amounts[i-v], amounts[i])
		}
	}

	res := amounts[len(amounts)-1]
	if res > amount {
		return -1
	}
	return res
}
