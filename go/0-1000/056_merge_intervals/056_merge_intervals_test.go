package mergeintervals

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		inters [][]int
		want   [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{[][]int{{1, 4}, {0, 0}}, [][]int{{0, 0}, {1, 4}}},
	}

	for _, tt := range tests {
		got := merge(tt.inters)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d: got %d, want %d", tt.inters, got, tt.want)
		}
	}
}
