package middleofthelinkedlist

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

func TestMiddle(t *testing.T) {
	tests := []struct {
		list []int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{4, 5, 6}},
		{[]int{1, 2}, []int{2}},
	}

	for _, tt := range tests {
		input := sliceToList(tt.list)
		output := middleNode(input)
		got := listToSlice(output)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d: got %d, want %d", tt.list, got, tt.want)
		}
	}
}
