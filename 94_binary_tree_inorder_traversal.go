package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	output := make([]int, 0)
	if root.Left != nil {
		output = append(output, inorderTraversal(root.Left)...)
	}
	output = append(output, root.Val)
	if root.Right != nil {
		output = append(output, inorderTraversal(root.Right)...)
	}

	return output
}

type stack struct {
	data []*TreeNode
}

func NewStack() stack {
	return stack{
		data: make([]*TreeNode, 0),
	}
}

func (s *stack) push(p *TreeNode) {
	s.data = append(s.data, p)
}

func (s *stack) pop() (*TreeNode, bool) {
	if len(s.data) == 0 {
		return nil, false
	}

	p := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return p, true
}

func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	s := NewStack()
	curr := root
	output := make([]int, 0)

	for curr != nil || len(s.data) > 0 {
		if curr == nil {
			curr, _ = s.pop()
			output = append(output, curr.Val)
			curr = curr.Right
		} else {
			s.push(curr)
			curr = curr.Left
		}
	}

	return output
}
