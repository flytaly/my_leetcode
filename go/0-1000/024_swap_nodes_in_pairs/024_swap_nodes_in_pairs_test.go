package swapnodesinpairs

import (
	. "my_leetcode"
	"reflect"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	tests := []struct {
		nodes []int
		want  []int
	}{
		{nodes: []int{1, 2, 3, 4}, want: []int{2, 1, 4, 3}},
		{nodes: []int{1}, want: []int{1}},
		{nodes: []int{1, 2, 3}, want: []int{2, 1, 3}},
		{nodes: []int{}, want: []int{}},
	}

	for _, tt := range tests {
		head := SliceToList(tt.nodes)
		got := ListToSlice(swapPairs(head))
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d: got %d, want %d", tt.nodes, got, tt.want)
		}
	}
}
