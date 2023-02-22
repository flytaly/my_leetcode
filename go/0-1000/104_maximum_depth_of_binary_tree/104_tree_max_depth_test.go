package maximum_depth_of_binary_tree

import "testing"

func TestMxDepth(t *testing.T) {
	tree := &TreeNode{
		3,
		&TreeNode{
			20,
			&TreeNode{7, nil, nil},
			&TreeNode{15, nil, nil},
		},
		&TreeNode{9, nil, nil},
	}

	got := maxDepth(tree)
	if got != 3 {
		t.Errorf("got %d, want %d", got, 3)
	}
}
