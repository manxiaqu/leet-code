package leetcode

import "fmt"

/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。



示例 1：


输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
示例 2：

输入：l1 = [0], l2 = [0]
输出：[0]
示例 3：

输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]


提示：

每个链表中的节点数在范围 [1, 100] 内
0 <= Node.val <= 9
题目数据保证列表表示的数字不含前导零
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) print() {
	p := l
	for p != nil {
		fmt.Printf("%d ", p.Val)
		p = p.Next
	}
	fmt.Println()
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0

	head := new(ListNode)
	l := head
	prev := l

	for l1 != nil && l2 != nil {
		l.Val = l1.Val + l2.Val + carry
		carry = 0

		if l.Val >= 10 {
			l.Val = l.Val - 10
			carry = 1
		}

		// step
		l1 = l1.Next
		l2 = l2.Next
		// todo: check if nil
		l.Next = new(ListNode)
		prev = l
		l = l.Next
	}

	if l1 == nil && l2 == nil && carry == 0 {
		prev.Next = nil
		return head
	}

	if l1 != nil {
		prev.Next = l1
		for l1 != nil && carry != 0 {
			l1.Val += carry
			carry = 0

			if l1.Val >= 10 {
				l1.Val = l1.Val - 10
				carry = 1
			}

			prev = l1
			l1 = l1.Next
		}
	}

	if l2 != nil {
		prev.Next = l2
		for l2 != nil && carry != 0 {
			l2.Val += carry
			carry = 0

			if l2.Val >= 10 {
				l2.Val = l2.Val - 10
				carry = 1
			}

			prev = l2
			l2 = l2.Next
		}
	}

	if carry == 1 {
		prev.Next = new(ListNode)
		prev.Next.Val = carry
	}

	return head
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	head := l1
	carry := 0
	prev := head

	for l1 != nil {
		if l2 == nil && carry == 0 {
			return head
		}

		if l2 != nil {
			l1.Val += l2.Val

			l2 = l2.Next
		}
		if carry > 0 {
			l1.Val += carry
		}

		carry = 0
		if l1.Val >= 10 {
			l1.Val = l1.Val - 10
			carry = 1
		}

		// step
		prev = l1
		l1 = l1.Next
	}

	// break only if l1 == nil
	if l2 != nil {
		prev.Next = l2
		for l2 != nil && carry != 0 {
			l2.Val += carry
			carry = 0

			if l2.Val >= 10 {
				l2.Val = l2.Val - 10
				carry = 1
			}

			prev = l2
			l2 = l2.Next
		}
	}

	if carry == 1 {
		prev.Next = new(ListNode)
		prev.Next.Val = carry
	}

	return head
}
