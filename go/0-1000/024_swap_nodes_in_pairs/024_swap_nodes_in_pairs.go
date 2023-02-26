package swapnodesinpairs

import (
	. "my_leetcode"
)

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev, next, swapped *ListNode = nil, nil, nil
	var resultHead = head.Next

	for node := head; node != nil; node = next {
		next = node.Next
		if prev != nil {
			node.Next = prev
			if swapped != nil {
				swapped.Next = node
			}
			swapped = prev
			prev = nil
			continue
		}

		prev = node
		prev.Next = nil
	}

	if prev != nil {
		swapped.Next = prev
	}

	return resultHead
}
