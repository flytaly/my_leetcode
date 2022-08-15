package maximum_depth_of_binary_tree

/*
104. Maximum Depth of Binary Tree

Given the root of a binary tree, return its maximum depth.
A binary tree's maximum depth is the number of nodes along
the longest path from the root node down to the farthest leaf node.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// non recursive
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 1
	queue := []*TreeNode{root, nil}

	for len(queue) > 1 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			depth++
			queue = append(queue, nil)
			continue
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return depth
}

/*
	// with recursion
	func maxDepth(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftDepth := maxDepth(root.Left)
		rightDepth := maxDepth(root.Right)
		if leftDepth > rightDepth {
			return leftDepth + 1
		} else {
			return rightDepth + 1
		}
	}
*/
