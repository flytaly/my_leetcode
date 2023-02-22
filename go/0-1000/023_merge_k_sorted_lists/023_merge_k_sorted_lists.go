package mergeksortedlists

import (
	"container/heap"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type PQueue []*ListNode

func (pq PQueue) Len() int { return len(pq) }

func (pq PQueue) Less(i, j int) bool { return (*pq[i]).Val < (*pq[j]).Val }

func (pq PQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PQueue) Push(x interface{}) { *pq = append(*pq, x.(*ListNode)) }

func (pq *PQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

// with priority queue
func mergeKLists(lists []*ListNode) *ListNode {
	pq := make(PQueue, 0)
	heap.Init(&pq)

	for _, ln := range lists {
		for ln != nil {
			heap.Push(&pq, ln)
			ln = ln.Next
		}
	}

	if pq.Len() == 0 {
		return nil
	}

	curr := heap.Pop(&pq).(*ListNode)
	head := curr

	for pq.Len() > 0 {
		ln := heap.Pop(&pq).(*ListNode)
		curr.Next = ln
		curr = curr.Next
	}
	curr.Next = nil
	return head
}

///================================================

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

func mergeKListsV1(lists []*ListNode) *ListNode {
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
