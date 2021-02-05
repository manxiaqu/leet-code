package leetcode

/*
给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。

进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？



示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
示例 3：

输入：nums1 = [0,0], nums2 = [0,0]
输出：0.00000
示例 4：

输入：nums1 = [], nums2 = [1]
输出：1.00000
示例 5：

输入：nums1 = [2], nums2 = []
输出：2.00000s
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return median(nums2)
	}
	if len(nums2) == 0 {
		return median(nums1)
	}

	total := len(nums1) + len(nums2)
	odd := total%2 != 0
	mid := total / 2

	midFirst := 0
	midSecond := 0

	nums1Index := 0
	nums2Index := 0

	current := 0
	currentIndex := 0
	for nums1Index < len(nums1) && nums2Index < len(nums2) {
		currentIndex = nums1Index + nums2Index
		if nums1[nums1Index] < nums2[nums2Index] {
			current = nums1[nums1Index]
			nums1Index++
		} else {
			current = nums2[nums2Index]
			nums2Index++
		}

		if currentIndex == mid-1 {
			midFirst = current
		}
		if currentIndex == mid {
			midSecond = current
			break
		}
	}

	if nums1Index >= len(nums1) {
		if mid-1 >= len(nums1)+nums2Index {
			midFirstIndex := mid - 1 - len(nums1)
			midFirst = nums2[midFirstIndex]
			midSecond = nums2[midFirstIndex+1]
		} else if mid >= len(nums1)+nums2Index {
			midSecond = nums2[mid-len(nums1)]
		}
	}

	if nums2Index >= len(nums2) {
		if mid-1 >= len(nums2)+nums1Index {
			midFirstIndex := mid - 1 - len(nums2)
			midFirst = nums1[midFirstIndex]
			midSecond = nums1[midFirstIndex+1]
		} else if mid >= len(nums2)+nums1Index {
			midSecond = nums1[mid-len(nums2)]
		}
	}

	if odd {
		return float64(midSecond)
	} else {
		return float64(midFirst+midSecond) / 2
	}
}

func median(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}

	mid := len(nums) / 2
	if len(nums)%2 == 0 {
		return float64(nums[mid]+nums[mid-1]) / 2
	} else {
		return float64(nums[mid])
	}
}

// func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
// 	odd := len(nums1)+len(nums2)%2 != 0
// 	mid := len(nums1) + len(nums2)/2

// 	midFirst := 0
// 	midSecond := 0

// 	index := 0
// 	i, j := 0, 0
// 	current := 0
// 	for {

// 		//

// 		if mid-1 == index {
// 			midFirst = current
// 		}
// 		if mid == index {
// 			midSecond = current
// 			break
// 		}

// 		index++
// 	}

// 	if odd {
// 		return float64(midSecond)
// 	} else {
// 		return float64(midFirst+midSecond) / 2
// 	}
// }
