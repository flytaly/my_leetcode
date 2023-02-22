package sametree

/*
100. Same Tree

Given the roots of two binary trees p and q, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// non-recursive
func isSameTree(p *TreeNode, q *TreeNode) bool {
	queue := [][]*TreeNode{{p, q}}

	for len(queue) > 0 {
		// if pop the last elements => dfs; if first element => bfs
		last := len(queue) - 1
		p := queue[last][0]
		q := queue[last][1]
		queue = queue[:last]

		if p == nil || q == nil {
			if p != nil || q != nil {
				return false
			}
			continue
		}

		if p.Val != q.Val {
			return false
		}

		queue = append(queue, []*TreeNode{p.Left, q.Left}, []*TreeNode{p.Right, q.Right})
	}

	return true
}

/*
	// recursion
	func isSameTree(p *TreeNode, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}

		if p == nil || q == nil {
			return false
		}

		if p.Val == q.Val {
			left := isSameTree(p.Left, q.Left)
			right := isSameTree(p.Right, q.Right)
			return left && right
		}

		return false
	}
*/
