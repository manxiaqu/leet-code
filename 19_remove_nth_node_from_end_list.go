package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	list := make([]*ListNode, 0)

	for head != nil {
		list = append(list, head)
		head = head.Next
	}

	index := len(list) - n
	if index == 0 {
		return list[index].Next
	}
	list[index-1].Next = list[index].Next

	return list[0]
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	list := head
	candidateBefore := head

	for list != nil {
		n--
		if n < -1 {
			candidateBefore = candidateBefore.Next
		}
		list = list.Next
	}

	if candidateBefore == head && n >= 0 {
		return head.Next
	}
	candidateBefore.Next = candidateBefore.Next.Next
	return head
}
