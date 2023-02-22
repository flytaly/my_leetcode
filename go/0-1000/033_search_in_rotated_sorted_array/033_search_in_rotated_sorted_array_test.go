package searchinrotatedsortedarray

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 6, 2},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 7, 3},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
		{[]int{1}, 1, 0},
		{[]int{3, 1}, 1, 1},
	}

	for _, tt := range tests {
		got := search(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d->%d: got %d, want %d", tt.nums, tt.target, got, tt.want)
		}
	}
}
