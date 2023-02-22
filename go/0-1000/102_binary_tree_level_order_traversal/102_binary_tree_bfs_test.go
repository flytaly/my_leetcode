package binarytreelevelordertraversal

import (
	"reflect"
	"testing"
)

type T = TreeNode

func Test_levelOrder(t *testing.T) {
	tests := []struct {
		name string
		root *T
		want [][]int
	}{
		{"example1",
			&T{Val: 3, Left: &T{Val: 9},
				Right: &T{
					Val: 20, Left: &T{Val: 15}, Right: &T{Val: 7},
				},
			}, [][]int{{3}, {9, 20}, {15, 7}}},
		{"example2", &T{Val: 1}, [][]int{{1}}},
		{"example3", nil, [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
