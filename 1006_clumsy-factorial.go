package leetcode

func clumsy(N int) int {
	subStack := make([]int, 0)
	curr := N
	count := 0

	for i := N - 1; i > 0; i-- {
		switch count & 3 {
		// *
		case 0:
			curr = curr * i
		// /
		case 1:
			curr = curr / i
		// +
		case 2:
			if len(subStack) > 0 {
				curr = subStack[0] - curr
				subStack = make([]int, 0)
			}
			curr = curr + i
		// -
		case 3:
			// no sub value
			// the first time
			if len(subStack) == 0 {
				subStack = append(subStack, curr)
			}
			curr = i
		}
		count++
	}

	if len(subStack) > 0 {
		curr = subStack[0] - curr
	}

	return curr
}
