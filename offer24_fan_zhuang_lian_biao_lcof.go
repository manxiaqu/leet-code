package leetcode

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev, curr, next *ListNode

	for head.Next != nil {
		curr = head
		next = head.Next

		// reverse
		curr.Next = prev
		head = next
		prev = curr
	}

	curr = head
	curr.Next = prev

	return curr
}
