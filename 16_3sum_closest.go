package leetcode

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	// first, second, third
	closet := nums[0] + nums[1] + nums[2]

	// 1. 设定first = 0， second=1, third=len(nums)-1，移动到最近的位置
	for first := 0; first < len(nums)-2; first++ {
		second, third := first+1, len(nums)-1

		var sum int
		for second < third {
			sum = nums[first] + nums[second] + nums[third]
			closetDiff := abs(closet - target)
			if abs(sum-target) < closetDiff {
				closet = sum
			}
			if sum > target {
				third--
			} else if sum < target {
				second++
			} else {
				return target
			}
		}
	}

	return closet
}

func abs(v int) int {
	if v > 0 {
		return v
	}

	return -v
}
