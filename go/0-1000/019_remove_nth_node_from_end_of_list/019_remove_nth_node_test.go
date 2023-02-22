package removenthnodefromendoflist

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

func TestRemove(t *testing.T) {
	tests := []struct {
		listVals []int
		num      int
		want     []int
	}{
		{listVals: []int{1, 2, 3, 4, 5}, num: 2, want: []int{1, 2, 3, 5}},
		{listVals: []int{1, 2, 3, 4, 5, 6, 7}, num: 4, want: []int{1, 2, 3, 5, 6, 7}},
		{listVals: []int{1}, num: 1, want: nil},
		{listVals: []int{1, 2}, num: 1, want: []int{1}},
	}

	for _, tt := range tests {
		list := sliceToList(tt.listVals)
		res := removeNthFromEnd(list, tt.num)
		got := listToSlice(res)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d->%d: got %d, want %d", tt.listVals, tt.num, got, tt.want)
		}
	}
}
