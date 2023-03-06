package findfirstandlastpos

import (
	"reflect"
	"testing"
)

func Test_searchRange(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{}, 0, []int{-1, -1}},
		{[]int{1}, 1, []int{0, 0}},
	}
	for _, tt := range tests {
		if got := searchRange(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("searchRange(%v, %v) = %v, want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}
