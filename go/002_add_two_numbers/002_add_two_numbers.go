package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	var head = &ListNode{}
	var node = head

	for {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		node.Val = sum % 10

		if l1 == nil && l2 == nil {
			break
		}

		node.Next = &ListNode{}
		node = node.Next
	}

	if carry > 0 {
		node.Next = &ListNode{Val: carry}
	}

	return head
}
