package mergeksortedlists

import (
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// returns index of the list with minimum value,
// returns -1 if there are no values anymore
func minIndex(ll []*ListNode) int {
	minVal, index := math.MaxInt, -1
	for i, ln := range ll {
		if ln != nil {
			if ln.Val < minVal {
				minVal = ln.Val
				index = i
			}
		}
	}
	return index
}

func mergeKLists(lists []*ListNode) *ListNode {
	minIdx := minIndex(lists)
	if minIdx == -1 {
		return nil
	}
	node := &ListNode{Val: lists[minIdx].Val}
	head := node
	lists[minIdx] = lists[minIdx].Next

	for {
		minIdx := minIndex(lists)
		if minIdx == -1 {
			return head
		}
		node.Next = &ListNode{Val: lists[minIdx].Val}
		node = node.Next
		lists[minIdx] = lists[minIdx].Next
	}
}
