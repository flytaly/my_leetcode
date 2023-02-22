package best_time_to_buy_and_sell_stock

/*
	121. Best Time to Buy and Sell Stock

	You are given an array prices where prices[i] is the price of a given stock on the ith day.

	You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

	Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0
*/

import "math"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	result := 0
	buyPrice := math.MaxInt

	for _, price := range prices {
		result = max(price-buyPrice, result)
		if price < buyPrice {
			buyPrice = price
		}
	}

	return result
}
