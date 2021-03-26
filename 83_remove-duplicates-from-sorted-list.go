package leetcode

func deleteDuplicates(head *ListNode) *ListNode {
	first := new(ListNode)
	first.Val = -101
	first.Next = head

	output := first

	for first.Next != nil {
		// delete first.next
		if first.Val == first.Next.Val {
			first.Next = first.Next.Next
			continue
		}

		first = first.Next
	}

	return output.Next
}
