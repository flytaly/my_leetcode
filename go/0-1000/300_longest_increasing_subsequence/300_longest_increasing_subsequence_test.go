package longestincreasingsubsequence

import (
	"reflect"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7, 7, 7}, 1},
		{[]int{0, 1, 10, 11, 12, 8, 9}, 5},
		{[]int{0, 1, 10, 11, 12, 8, 9, 11, 12, 9}, 6},
	}

	for _, tt := range tests {
		got := lengthOfLIS(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d: got %d, want %d", tt.nums, got, tt.want)
		}
	}
}
