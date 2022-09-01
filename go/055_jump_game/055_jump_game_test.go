package jumpgame

import "testing"

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"example_1", args{[]int{2, 3, 1, 1, 4}}, true},
		{"example_2", args{[]int{3, 2, 1, 0, 4}}, false},
		{"example_3", args{[]int{0, 1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
