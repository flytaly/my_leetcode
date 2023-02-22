/* 206. Reverse Linked List */

package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// iterative
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for ; head != nil; head = head.Next {
		prev = &ListNode{Val: head.Val, Next: prev}
	}

	return prev
}

// iterative: in-place
func reverseListInPlace(head *ListNode) *ListNode {
	var prev *ListNode
	var next *ListNode

	for current := head; current != nil; current = next {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}
