package subtreeofanothertree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}

	return isSameTree(t1.Left, t2.Left) && isSameTree(t1.Right, t2.Right)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if isSameTree(root, subRoot) {
		return true
	}
	if root == nil || subRoot == nil {
		return false
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

/* func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	bfsQueue := []*TreeNode{root}

	for len(bfsQueue) > 0 {
		node := bfsQueue[0]
		bfsQueue = bfsQueue[1:]

		if isSameTree(node, subRoot) {
			return true
		}

		if node != nil {
			bfsQueue = append(bfsQueue, node.Right, node.Left)
		}
	}

	return false
} */
