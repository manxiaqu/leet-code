package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	target := 0
	output := make([][]int, 0)

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if target = nums[i]; target > 0 {
			break
		}

		l, r := i+1, len(nums)-1
		for l < r {
			if nums[l]+nums[r]+target < 0 {
				l++
			} else if nums[l]+nums[r]+target > 0 {
				r--
			} else {
				output = append(output, []int{target, nums[l], nums[r]})

				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
		}
	}

	return output
}
