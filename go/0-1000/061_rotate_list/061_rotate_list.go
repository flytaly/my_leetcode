package rotatelist

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var length = 0
	for node := head; node != nil; node = node.Next {
		length += 1
	}
	shift := (length + k) % length
	headIndex := (length - shift) % length

	if headIndex == 0 {
		return head
	}

	var newHead *ListNode
	var node = head
	for i := 0; i < length; i += 1 {
		next := node.Next

		switch i {
		case headIndex:
			newHead = node
		case headIndex - 1:
			node.Next = nil
		}

		if i == length-1 {
			node.Next = head
		}
		node = next
	}

	return newHead
}
