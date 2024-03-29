package my_leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func SliceToList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := ListNode{Val: vals[0], Next: nil}
	last := &head

	for _, v := range vals[1:] {
		last.Next = &ListNode{Val: v, Next: nil}
		last = last.Next
	}

	return &head
}

func ListToSlice(head *ListNode) (res []int) {
	if head == nil {
		return []int{}
	}

	for ; head != nil; head = head.Next {
		res = append(res, head.Val)
	}
	return res
}
