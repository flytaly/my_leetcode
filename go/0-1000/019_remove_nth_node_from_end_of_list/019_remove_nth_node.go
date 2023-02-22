package removenthnodefromendoflist

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var del *ListNode

	delIdx := -n
	for curr := head; curr != nil; curr = curr.Next {
		if delIdx == 0 {
			del = head
		}
		if delIdx > 0 {
			del = del.Next
		}
		delIdx++
	}

	if del != nil && del.Next != nil {
		del.Next = del.Next.Next
		return head
	}

	if delIdx == 0 && head != nil {
		head = head.Next
	}

	return head
}
