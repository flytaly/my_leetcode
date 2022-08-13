package insertinterval

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		inters   [][]int
		newInter []int
		want     [][]int
	}{
		{[][]int{{1, 3}, {6, 9}}, []int{2, 5}, [][]int{{1, 5}, {6, 9}}},
		{[][]int{{1, 3}, {6, 9}}, []int{2, 20}, [][]int{{1, 20}}},
		{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}, {20, 30}}, []int{4, 8}, [][]int{{1, 2}, {3, 10}, {12, 16}, {20, 30}}},
	}

	for _, tt := range tests {
		got := insert(tt.inters, tt.newInter)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d -> %d: got %d, want %d", tt.newInter, tt.inters, got, tt.want)
		}
	}
}
