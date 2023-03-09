package removeelement

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		args args
		want int
	}{
		{args{[]int{3, 2, 2, 3}, 3}, 2},
		{args{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2}, 5},
	}
	for _, tt := range tests {
		if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
			t.Errorf("removeElement() = %v, want %v", got, tt.want)
		}
	}
}
