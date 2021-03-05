package leetcode

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	current, prevBefore, prev := 0, 1, 1
	for i := 3; i <= n; i++ {
		current = prevBefore + prev
		prevBefore, prev = prev, current
	}

	return current
}
