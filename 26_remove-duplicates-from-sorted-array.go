package leetcode

func removeDuplicates3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	index := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			index++
			nums[index] = nums[i]
		}
	}

	return index + 1
}
