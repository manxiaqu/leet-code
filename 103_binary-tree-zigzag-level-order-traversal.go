package leetcode

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	output := make([][]int, 0)
	levelNodes := []*TreeNode{root}
	level := 1

	for len(levelNodes) > 0 {
		tmpLevelNodes := make([]*TreeNode, 0)
		num := len(levelNodes)
		tmpOutput := make([]int, num)

		for i := 0; i < num; i++ {
			node := levelNodes[i]
			if level%2 == 1 {
				tmpOutput[i] = node.Val
			} else {
				tmpOutput[i] = levelNodes[num-1-i].Val
			}

			if node.Left != nil {
				tmpLevelNodes = append(tmpLevelNodes, node.Left)
			}
			if node.Right != nil {
				tmpLevelNodes = append(tmpLevelNodes, node.Right)
			}
		}

		output = append(output, tmpOutput)
		levelNodes = tmpLevelNodes
		level++
	}

	return output
}
