package binarytreemaximumpathsum

import "testing"

type T = TreeNode

func Test_maxPathSum(t *testing.T) {
	tests := []struct {
		name string
		root *T
		want int
	}{
		{"example1", &T{Val: 1, Left: &T{Val: 2}, Right: &T{Val: 3}},
			6},
		{"example2",
			&T{
				Val:  -10,
				Left: &T{Val: 9},
				Right: &T{
					Val:   20,
					Left:  &T{Val: 15},
					Right: &T{Val: 7},
				},
			},
			42},
		{"example3", &T{Val: -3}, -3},
		{"example4", &T{Val: 2, Left: &T{Val: -1}}, 2},
		{"example5", &T{Val: 1, Left: &T{Val: -2}, Right: &T{Val: 3}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
