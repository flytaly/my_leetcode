package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

// recursive
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}

func mergeTwoListsV1(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var list *ListNode

	if list1.Val < list2.Val {
		list = &ListNode{Val: list1.Val}
		list1 = list1.Next
	} else {
		list = &ListNode{Val: list2.Val}
		list2 = list2.Next
	}

	head := list

	for list1 != nil && list2 != nil {
		list.Next = &ListNode{}
		list = list.Next
		if list1.Val < list2.Val {
			list.Val = list1.Val
			list1 = list1.Next
		} else {
			list.Val = list2.Val
			list2 = list2.Next
		}
	}

	if list1 == nil && list2 != nil {
		list.Next = list2
	}
	if list2 == nil && list1 != nil {
		list.Next = list1
	}

	return head
}
