package leetcode

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	// step 1
	slow := head
	// step 2
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			q := head
			for q != slow {
				q = q.Next
				slow = slow.Next
			}
			return q
		}
	}
	return nil
}
