package reorderlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func reverseList(head *ListNode) *ListNode {
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

func reorderList(head *ListNode) {
	tail := middleNode(head)
	next := tail
	tail = tail.Next
	next.Next = nil
	tail = reverseList(tail)
	for tail != nil {
		next := head.Next
		head.Next = tail
		tail = tail.Next
		head.Next.Next = next
		head = next
	}
}
