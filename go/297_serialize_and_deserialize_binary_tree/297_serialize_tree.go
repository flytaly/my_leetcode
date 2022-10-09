package serializeanddeserializebinarytree

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	result := []string{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			result = append(result, "null")
			continue
		}
		result = append(result, strconv.Itoa(node.Val))
		queue = append(queue, node.Left, node.Right)
	}

	length := len(result)
	for i := len(result) - 1; i >= 0; i-- {
		if result[i] != "null" {
			length = i
			break
		}
	}
	return strings.Join(result[0:length+1], ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	root := &TreeNode{}

	values := strings.Split(data, ",")
	queue := []*TreeNode{root}
	var err error
	root.Val, err = strconv.Atoi(values[0])
	if err != nil {
		return nil
	}

	for i := 1; i < len(values); i++ {
		node := queue[0]
		queue = queue[1:]
		lNum, lErr := strconv.Atoi(values[i])
		if lErr == nil {
			node.Left = &TreeNode{Val: lNum}
			queue = append(queue, node.Left)
		}
		if i+1 < len(values) {
			rNum, rErr := strconv.Atoi(values[i+1])
			if rErr == nil {
				node.Right = &TreeNode{Val: rNum}
				queue = append(queue, node.Right)
			}
		}
		i += 1
	}
	return root
}
