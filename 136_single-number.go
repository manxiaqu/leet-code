package leetcode

func singleNumber_1(nums []int) int {
	output := 0
	for _, v := range nums {
		output ^= v
	}

	return output
}
