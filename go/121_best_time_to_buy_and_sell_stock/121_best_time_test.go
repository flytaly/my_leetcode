package best_time_to_buy_and_sell_stock

import (
	"reflect"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		// Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		// In this case, no transactions are done and the max profit = 0.
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, tt := range tests {
		got := maxProfit(tt.prices)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d: got %d, want %d", tt.prices, got, tt.want)
		}
	}
}
