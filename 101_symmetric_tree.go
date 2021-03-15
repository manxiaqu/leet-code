package leetcode

func isSymmetric(root *TreeNode) bool {
	return root == nil || isTwoTreeSymmetric(root.Left, root.Right)
}

func isTwoTreeSymmetric(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == nil && right == nil
	}
	if left.Val != right.Val {
		return false
	}
	return isTwoTreeSymmetric(left.Left, right.Right) && isTwoTreeSymmetric(left.Right, right.Left)
}

// 中左右
// 中右左
// 得到的结果一致
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// 遍历左子树
	// 遍历右子树
	left := root.Left
	right := root.Right

	leftStack := make([]*TreeNode, 0)
	rightStack := make([]*TreeNode, 0)

	for {
		if (left == nil && right != nil) || (right == nil && left != nil) {
			return false
		}
		if left == nil {
			leftLen, rightLen := len(leftStack), len(rightStack)
			if leftLen != rightLen {
				return false
			}
			if rightLen == 0 {
				break
			}
			// pop from stack to check
			left = leftStack[leftLen-1].Right
			leftStack = leftStack[:leftLen-1]

			right = rightStack[rightLen-1].Left
			rightStack = rightStack[:rightLen-1]

			continue
		}

		if left.Val != right.Val {
			return false
		}

		leftStack = append(leftStack, left)
		rightStack = append(rightStack, right)

		// get next left
		left = left.Left

		// get next right
		right = right.Right
	}

	return true
}
