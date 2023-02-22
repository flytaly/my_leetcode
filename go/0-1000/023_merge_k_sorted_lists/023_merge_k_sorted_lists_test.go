package mergeksortedlists

import (
	"reflect"
	"testing"
)

func sliceToList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := ListNode{Val: vals[0], Next: nil}
	last := &head

	for _, v := range vals[1:] {
		last.Next = &ListNode{Val: v, Next: nil}
		last = last.Next
	}

	return &head
}

func listToSlice(head *ListNode) (res []int) {
	for ; head != nil; head = head.Next {
		res = append(res, head.Val)
	}
	return res
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		lists [][]int
		want  []int
	}{
		{lists: [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, want: []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{lists: [][]int{}, want: nil},
		{lists: [][]int{{}}, want: nil},
		{lists: [][]int{{-2, -1, -1, -1}, {}}, want: []int{-2, -1, -1, -1}},
	}

	for _, tt := range tests {
		ll := []*ListNode{}
		for _, l := range tt.lists {
			ll = append(ll, sliceToList(l))
		}
		res := mergeKLists(ll)
		got := listToSlice(res)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d: got %d, want %d", tt.lists, got, tt.want)
		}
	}
}
