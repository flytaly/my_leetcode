package linkedlistcycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	tortoise := head.Next
	hare := head.Next.Next

	for hare != nil {
		if hare == tortoise {
			return true
		}
		if hare.Next == nil {
			return false
		}
		tortoise = tortoise.Next
		hare = hare.Next.Next
	}

	return false
}
