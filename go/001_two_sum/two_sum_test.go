package twosums

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{11, 15, 2, 7}, 9, []int{2, 3}},
		{[]int{2, 5, 5, 11}, 10, []int{1, 2}},
		{[]int{2, 5, 11, 7, 3}, 10, []int{3, 4}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
	}

	for _, tt := range tests {
		got := twoSum(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d->%d: got %d, want %d", tt.nums, tt.target, got, tt.want)
		}
	}
}
