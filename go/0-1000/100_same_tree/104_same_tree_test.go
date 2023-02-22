package sametree

import "testing"

type TN = TreeNode

func TestMxDepth(t *testing.T) {
	tests := []struct {
		tree1  *TN
		tree2  *TN
		isSame bool
	}{
		{
			&TN{1, &TN{2, nil, nil}, &TN{3, nil, nil}},
			&TN{1, &TN{2, nil, nil}, &TN{3, nil, nil}},
			true,
		},
		{
			&TN{1, &TN{2, nil, nil}, &TN{1, nil, nil}},
			&TN{1, &TN{1, nil, nil}, &TN{2, nil, nil}},
			false,
		},
		{
			&TN{1, &TN{2, nil, nil}, &TN{3, nil, nil}},
			&TN{1, &TN{2, nil, nil}, nil},
			false,
		},
	}

	for i, tt := range tests {
		got := isSameTree(tt.tree1, tt.tree2)
		if got != tt.isSame {
			t.Errorf("%d: got %t, want %t", i, got, tt.isSame)
		}
	}
}
