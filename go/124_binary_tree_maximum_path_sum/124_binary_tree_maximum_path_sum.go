package binarytreemaximumpathsum

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxOf(vars ...int) int {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}

func dfs(node *TreeNode) (openSum int, maxSum int) {
	if node == nil {
		return 0, math.MinInt
	}

	lSum, lMaxSum := dfs(node.Left)
	rSum, rMaxSum := dfs(node.Right)

	openSum = MaxOf(node.Val+rSum, node.Val+lSum, node.Val)
	maxSum = MaxOf(lMaxSum, rMaxSum, openSum, lSum+node.Val+rSum)
	return openSum, maxSum
}

func maxPathSum(root *TreeNode) int {
	_, maxSum := dfs(root)
	return maxSum
}
