package leetcode

func climbStairs(n int) int {
	// f(n) = f(n-1) + f(n-2)
	// n=1, 1
	// n=2, 2
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	first, second := 1, 2
	for i := 3; i < n; i++ {
		first, second = second, first+second
	}
	return first + second
}
