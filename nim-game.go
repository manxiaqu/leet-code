package leetcode

func canWinNim(n int) bool {
	// n%4 != 0
	return n&3 != 0
}
