package threeSum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{-1, 2, 1, -4}, 1, 2},
		{[]int{0, 0, 0}, 1, 0},
	}

	for _, tt := range tests {
		got := threeSumClosest(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("(%d, %d) => %d, want %d", tt.nums, tt.target, got, tt.want)
		}
	}
}
