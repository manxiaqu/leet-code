package leetcode

func isPowerOfFour(n int) bool {
	if n == 0 {
		return false
	}

	// is power of 2
	if n&(n-1) == 0 {
		count := 0
		for n > 0 {
			n = n >> 1
			count++
		}

		if count&1 == 1 {
			return true
		}
	}

	return false
}
