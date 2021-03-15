package leetcode

func isValidBST(root *TreeNode) bool {
	check := root
	stack := make([]*TreeNode, 0)

	first := true
	prev := 0

	for check != nil || len(stack) != 0 {
		if check == nil {
			// pop from stack
			// should only check right node
			check = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if first {
				first = false
			} else if prev >= check.Val {
				return false
			}
			prev = check.Val
			check = check.Right
		} else {
			stack = append(stack, check)

			check = check.Left
		}
	}

	return true
}

/*
利用中序遍历，左中右
访问的节点的值是一个从小到大的序列

先遍历根节点，
根节点左子节点不为空，则将当前节点押入栈，检查左子节点
否则 从栈中取出元素，检查右子节点
*/
