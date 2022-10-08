package binarytreelevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// bfs
func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	level := []*TreeNode{root}
	nextLevel := []*TreeNode{}
	result = append(result, []int{})
	for len(level) > 0 {
		node := level[0]
		result[len(result)-1] = append(result[len(result)-1], node.Val)
		if node.Left != nil {
			nextLevel = append(nextLevel, node.Left)
		}
		if node.Right != nil {
			nextLevel = append(nextLevel, node.Right)
		}
		if len(level) > 1 {
			level = level[1:]
			continue
		}
		if len(nextLevel) == 0 {
			return result
		}
		level = nextLevel
		nextLevel = []*TreeNode{}
		result = append(result, []int{})
	}

	return result
}
