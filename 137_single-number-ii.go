package leetcode

import "sort"

func singleNumber(nums []int) int {
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if (i-1 < 0 || nums[i-1] != nums[i]) && (i >= len(nums)-1 || nums[i+1] != nums[i]) {
			return nums[i]
		}
	}

	return 0
}

func singleNumber2(nums []int) int {
	for i := 0; i < len(nums); i = i + 3 {
		if i == len(nums)-1 {
			return nums[i]
		}
		count := 1
		for base, k := i, i+1; k < len(nums); k++ {
			if nums[k] == nums[base] {
				// swap element
				tmp := nums[base+1]
				nums[base+1] = nums[k]
				nums[k] = tmp
				count++
				base++
			}

			if count == 3 {
				break
			}
		}

		if count == 1 {
			return nums[i]
		}
	}

	return 0
}

func singleNumber3(nums []int) int {
	seenOnce, seenTwice := 0, 0

	for _, num := range nums {
		// seen once存储出现了1次的结果
		seenOnce = ^seenTwice & (seenOnce ^ num)
		// seen twice存储出现了2次的结果
		seenTwice = ^seenOnce & (seenTwice ^ num)
	}
	return seenOnce
}
