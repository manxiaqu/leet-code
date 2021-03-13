package leetcode

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	}
	if len(cost) == 1 {
		return cost[0]
	}
	// mina 代表从a踏出最小花费
	mina := cost[0]
	// minb 代表从b踏出最小花费
	minb := cost[1]

	for i := 2; i < len(cost); i++ {
		mina, minb = minb, min(mina+cost[i], minb+cost[i])
	}

	return min(mina, minb)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
