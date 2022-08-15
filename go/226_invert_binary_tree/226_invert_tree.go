package main

/*
226. Invert Binary Tree

Given the root of a binary tree, invert the tree, and return its root.

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursion
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{
		root.Val,
		invertTree(root.Right),
		invertTree(root.Left),
	}
}
