package clonegraph

import (
	"reflect"
	"testing"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Test case format:
//
// For simplicity, each node's value is the same as the node's index (1-indexed). For example, the first node with val == 1, the second node with val == 2, and so on. The graph is represented in the test case using an adjacency list.
//
// An adjacency list is a collection of unordered lists used to represent a finite graph. Each list describes the set of neighbors of a node in the graph.

func sliceToGraph(s [][]int) *Node {
	if len(s) == 0 {
		return nil
	}
	nodes := []*Node{}

	for i := range s {
		nodes = append(nodes, &Node{Val: i + 1})
	}

	for i, adjacent := range s {
		for _, v := range adjacent {
			nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[v-1])
		}
	}

	return nodes[0]
}

func graphToSlice(node *Node) [][]int {
	if node == nil {
		return [][]int{}
	}
	resultMap := map[int][]int{}
	addList := []*Node{node}
	maxVal := 0
	for len(addList) > 0 {
		nextNode := addList[0]
		addList = addList[1:]
		resultMap[nextNode.Val] = []int{}
		maxVal = max(maxVal, nextNode.Val)
		for _, n := range nextNode.Neighbors {
			resultMap[nextNode.Val] = append(resultMap[nextNode.Val], n.Val)
			if _, ok := resultMap[n.Val]; ok != true {
				addList = append(addList, n)
			}
		}
	}

	result := make([][]int, maxVal)
	for i := 0; i < maxVal; i++ {
		result[i] = resultMap[i+1]
	}

	return result
}

func Test_cloneGraph(t *testing.T) {
	tests := []struct {
		name string
		node [][]int
		want [][]int
	}{
		{
			"example_1",
			[][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
			[][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
		},
		{"example_2", [][]int{{}}, [][]int{{}}},
		{"example_3", [][]int{}, [][]int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := sliceToGraph(tt.node)
			gotGraph := cloneGraph(node)
			got := graphToSlice(gotGraph)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cloneGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}
