package addtwonumbers

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
		a    []int
		b    []int
		want []int
	}{
		{[]int{0}, []int{0}, []int{0}},
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{2, 4, 8}, []int{5, 6, 4}, []int{7, 0, 3, 1}},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}

	for _, tt := range tests {
		inputA := sliceToList(tt.a)
		inputB := sliceToList(tt.b)
		got := listToSlice(addTwoNumbers(inputA, inputB))

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d + %d => got %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}
