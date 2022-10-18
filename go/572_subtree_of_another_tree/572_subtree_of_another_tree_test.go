package subtreeofanothertree

import "testing"

type T = TreeNode

func Test_isSubtree(t *testing.T) {
	type args struct {
		root    *TreeNode
		subRoot *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"example1",
			args{
				root: &T{
					Val:   4,
					Left:  &T{Val: 4, Left: &T{Val: 1}, Right: &T{Val: 2}},
					Right: &T{Val: 5},
				},
				subRoot: &T{Val: 4, Left: &T{Val: 1}, Right: &T{Val: 2}},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubtree(tt.args.root, tt.args.subRoot); got != tt.want {
				t.Errorf("isSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
