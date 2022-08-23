package coinchange

import (
	"reflect"
	"testing"
)

func TestCoinChange(t *testing.T) {
	tests := []struct {
		nums   []int
		amount int
		want   int
	}{
		{[]int{1, 2, 5, 33}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
		{[]int{2, 5, 10, 1}, 27, 4},
		{[]int{83, 186, 408, 419}, 6249, 20},
	}

	for _, tt := range tests {
		got := coinChange(tt.nums, tt.amount)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d->%d: got %d, want %d", tt.nums, tt.amount, got, tt.want)
		}
	}
}
