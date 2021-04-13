package leetcode

import "math"

func minDiffInBST(root *TreeNode) int {
	// 左边的数据都小于root
	// 右边的数据都大于root
	// 使用左中右的遍历顺序完成
	sorted := getArray(root)

	diff := sorted[1] - sorted[0]

	for i := 2; i < len(sorted); i++ {
		if sorted[i]-sorted[i-1] < diff {
			diff = sorted[i] - sorted[i-1]
		}
	}

	return diff
}

func getArray(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	output := make([]int, 0)

	if root.Left != nil {
		output = append(output, getArray(root.Left)...)
	}

	output = append(output, root.Val)

	if root.Right != nil {
		output = append(output, getArray(root.Right)...)
	}

	return output
}

func minDiffInBST2(root *TreeNode) int {
	// 左边的数据都小于root
	// 右边的数据都大于root
	// 使用左中右的遍历顺序完成
	curr := root
	stack := make([]*TreeNode, 0)
	var prev *TreeNode
	diff := math.MaxInt32

	for curr != nil || len(stack) != 0 {
		if curr != nil {
			stack = append(stack, curr)

			curr = curr.Left
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if prev != nil && curr.Val-prev.Val < diff {
				diff = curr.Val - prev.Val
			}

			prev = curr

			curr = curr.Right
		}
	}

	return diff
}
