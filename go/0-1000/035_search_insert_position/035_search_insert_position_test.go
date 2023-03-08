package searchinsert

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		args args
		want int
	}{
		{args{[]int{1, 3, 5, 6}, 5}, 2},
		{args{[]int{1, 3, 5, 6}, 2}, 1},
		{args{[]int{1, 3, 5, 5, 5, 8}, 7}, 5},
	}
	for _, tt := range tests {
		if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
			t.Errorf("searchInsert(%v, %v) = %v, want %v", tt.args.nums, tt.args.target, got, tt.want)
		}
	}
}
