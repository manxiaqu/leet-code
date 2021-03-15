package leetcode

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	output := make([][]int, 0)

	levelNodes := []*TreeNode{root}

	for len(levelNodes) > 0 {
		tmpLevelNodes := make([]*TreeNode, 0)
		tmpOutput := make([]int, 0)

		for _, node := range levelNodes {
			tmpOutput = append(tmpOutput, node.Val)

			if node.Left != nil {
				tmpLevelNodes = append(tmpLevelNodes, node.Left)
			}
			if node.Right != nil {
				tmpLevelNodes = append(tmpLevelNodes, node.Right)
			}
		}

		output = append(output, tmpOutput)
		levelNodes = tmpLevelNodes
	}

	return output
}
