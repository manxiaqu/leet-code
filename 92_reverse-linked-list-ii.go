package leetcode

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	var leftPos *ListNode
	check := head

	for check != nil {
		// 找到第一个位置
		right--
		if left--; left > 0 {
			leftPos = check
			check = check.Next
			continue
		}

		// 开始翻转
		prevBase := check
		prev := check
		check = check.Next

		for check != nil && right > 0 {
			tmp := check.Next
			check.Next = prev
			prev = check
			check = tmp

			right--
		}

		prevBase.Next = check
		if leftPos != nil {
			leftPos.Next = prev
			return head
		}

		return prev
	}

	return nil
}
