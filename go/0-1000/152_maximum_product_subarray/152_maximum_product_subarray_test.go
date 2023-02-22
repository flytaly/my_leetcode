package maximumproductsubarray

import (
	"reflect"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
		{[]int{-2, 3, -2, 4}, 48},
		{[]int{-2, 3, -2, 4, -4}, 96},
		{[]int{-4, -4, -2, -3}, 96},
		{[]int{-4, 0, -4, -2, -3}, 8},
		{[]int{-4}, -4},
		{[]int{2, -5, -2, -4, 3}, 24},
	}

	for _, tt := range tests {
		got := maxProduct(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d got %d, want %d", tt.nums, got, tt.want)
		}
	}
}
