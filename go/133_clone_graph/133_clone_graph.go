package clonegraph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}

	head := &Node{Val: node.Val}

	nodeMap := map[int]*Node{node.Val: head}
	nodeList := []*Node{node}

	for len(nodeList) > 0 {
		current := nodeList[len(nodeList)-1]
		nodeList = nodeList[:len(nodeList)-1]
		neighbors := []*Node{}
		for _, n := range current.Neighbors {
			if _, exists := nodeMap[n.Val]; exists == true {
				neighbors = append(neighbors, nodeMap[n.Val])
				continue
			}
			nodeMap[n.Val] = &Node{Val: n.Val}
			neighbors = append(neighbors, nodeMap[n.Val])
			nodeList = append(nodeList, n)
		}
		nodeMap[current.Val].Neighbors = neighbors
	}

	return head
}
