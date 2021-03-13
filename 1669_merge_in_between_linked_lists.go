package leetcode

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	head := list1
	// delete index [a,b]
	// get to a point
	var toRemovePrev, toRemoveAfter *ListNode
	index := 0
	for {
		if index < a {
			toRemovePrev = list1
		}
		if index == b {
			toRemoveAfter = list1.Next
			break
		}
		list1 = list1.Next
		index++
	}
	toRemovePrev.Next = list2

	if list2 != nil {
		for list2.Next != nil {
			list2 = list2.Next
		}
		list2.Next = toRemoveAfter
	}

	return head
}
