package nonoverlapping_intervals

import (
	"reflect"
	"testing"
)

func TestErase(t *testing.T) {
	tests := []struct {
		inters [][]int
		want   int
	}{
		{[][]int{{1, 3}, {1, 2}, {2, 3}, {3, 4}}, 1}, // [1,3] can be removed
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, 1}, // [1,3] can be removed
		{[][]int{{1, 2}, {1, 2}, {1, 2}}, 2},         // You need to remove two [1,2] to make the rest of the intervals non-overlapping.
		{[][]int{{1, 2}, {2, 3}}, 0},                 // You don't need to remove any of the intervals since they're already non-overlapping.
		{[][]int{{1, 2}}, 0},
		{[][]int{{0, 1}, {3, 4}, {1, 2}}, 0},
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {-100, -2}, {5, 7}}, 0},
		{[][]int{{-52, 31}, {-73, -26}, {82, 97}, {-65, -11}, {-62, -49}, {95, 99}, {58, 95}, {-31, 49}, {66, 98}, {-63, 2}, {30, 47}, {-40, -26}}, 7},
	}

	for _, tt := range tests {
		got := eraseOverlapIntervals(tt.inters)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d: got %d, want %d", tt.inters, got, tt.want)
		}
	}
}
