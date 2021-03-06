package leetcode

/*
给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。数字 x 的下一个更大的元素是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。

示例 1:

输入: [1,2,1]
输出: [2,-1,2]
解释: 第一个 1 的下一个更大的数是 2；
数字 2 找不到下一个更大的数；
第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
*/

func nextGreaterElements(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	output := make([]int, len(nums))
	// store index
	stack := make([]int, 1)
	stack[0] = 0

	for i := 1; i < len(nums); i++ {
		j := len(stack) - 1
		if len(stack) == 0 || nums[i] <= nums[stack[j]] {
			stack = append(stack, i)
			continue
		}
		// check curr and stack top
		for ; j >= 0 && nums[i] > nums[stack[j]]; j-- {
			output[stack[j]] = nums[i]
		}

		stack = stack[:j+1]
		stack = append(stack, i)
	}

	biggestIndex := stack[0]
	for i := 0; i <= biggestIndex; i++ {
		// check curr and stack top
		j := len(stack) - 1
		for ; j >= 0 && nums[i] > nums[stack[j]]; j-- {
			output[stack[j]] = nums[i]
		}

		stack = stack[:j+1]
	}

	for i := 0; i < len(stack); i++ {
		output[stack[i]] = -1
	}

	return output
}
