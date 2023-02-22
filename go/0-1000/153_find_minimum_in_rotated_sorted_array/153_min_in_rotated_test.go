package findminimuminrotatedsortedarray

import (
	"reflect"
	"testing"
)

func TestFindMin(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{11, 13, 15, 17}, 11},
		{[]int{10, 6, 7, 8, 9}, 6},
		{[]int{1, 2, 3}, 1},
		{[]int{2, 1}, 1},
	}

	for _, tt := range tests {
		got := findMin(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d: got %d, want %d", tt.nums, got, tt.want)
		}
	}
}
