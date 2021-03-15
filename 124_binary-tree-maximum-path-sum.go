package leetcode

import "math"

func maxPathSum(root *TreeNode) int {
	path := paths(root)

	return max(path[0], path[1])
}

func paths(node *TreeNode) [2]int {
	if node == nil {
		return [2]int{}
	}

	biggest := node.Val
	unchangeBiggest := node.Val
	tmp := node.Val

	if node.Left != nil {
		leftPaths := paths(node.Left)

		// get biggest path from three path
		// leftPaths[0] max path from this
		// leftPaths[1] max 独立path
		if leftPaths[0] > 0 {
			biggest += leftPaths[0]
		}

		unchangeBiggest = max(unchangeBiggest, leftPaths[1])
		unchangeBiggest = max(unchangeBiggest, leftPaths[0])
		tmp += leftPaths[0]
	}

	if node.Right != nil {
		rightPaths := paths(node.Right)

		if rightPaths[0] > 0 {
			biggest = max(biggest, node.Val+rightPaths[0])
		}

		unchangeBiggest = max(unchangeBiggest, rightPaths[1])
		unchangeBiggest = max(unchangeBiggest, rightPaths[0])
		tmp += rightPaths[0]
	}

	unchangeBiggest = max(unchangeBiggest, tmp)

	return [2]int{biggest, unchangeBiggest}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var ret = math.MinInt32

func maxPathSum2(root *TreeNode) int {
	getMax(root)

	return ret
}

func getMax(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := max(0, getMax(node.Left))
	right := max(0, getMax(node.Right))

	ret = max(ret, node.Val+left+right)

	return max(node.Val+left, node.Val+right)
}
