package house_robber_II

import (
	"reflect"
	"testing"
)

func TestRob(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, 2}, 3},
		{[]int{1, 2, 3, 1}, 4},
		{[]int{1, 2, 3}, 3},
		{[]int{1}, 1},
	}

	for _, tt := range tests {
		got := rob(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d: got %d, want %d", tt.nums, got, tt.want)
		}
	}
}
