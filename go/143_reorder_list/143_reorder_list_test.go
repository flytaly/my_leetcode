package reorderlist

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

func TestReorderList(t *testing.T) {
	tests := []struct {
		list []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 4, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 5, 2, 4, 3}},
	}

	for _, tt := range tests {
		input := sliceToList(tt.list)
		reorderList(input)
		got := listToSlice(input)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d: got %d, want %d", tt.list, got, tt.want)
		}
	}
}
