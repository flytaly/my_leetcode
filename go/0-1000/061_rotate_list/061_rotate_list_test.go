package rotatelist

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

func Test_rotateRight(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				head: sliceToList([]int{1, 2, 3, 4, 5}),
				k:    2,
			},
			want: []int{4, 5, 1, 2, 3},
		},
		{
			name: "2",
			args: args{
				head: sliceToList([]int{0, 1, 2}),
				k:    4,
			},
			want: []int{2, 0, 1},
		},
		{
			name: "3",
			args: args{
				head: sliceToList([]int{1, 2}),
				k:    0,
			},
			want: []int{1, 2},
		},
		{
			name: "4",
			args: args{
				head: sliceToList([]int{1, 2}),
				k:    2,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		got := listToSlice(rotateRight(tt.args.head, tt.args.k))
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("rotateRight() = %v, want %v", got, tt.want)
		}
	}
}
