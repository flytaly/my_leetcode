package reverselinkedlist

import (
	"reflect"
	"testing"
)

func sliceToList(vals []int) *ListNode {
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
		nodeValues []int
		want       []int
	}{
		{nodeValues: []int{1, 2, 3, 4, 5}, want: []int{5, 4, 3, 2, 1}},
	}

	for _, tt := range tests {
		head := sliceToList(tt.nodeValues)
		got := listToSlice(reverseList(head))
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d: got %d, want %d", tt.nodeValues, got, tt.want)
		}
	}
}
