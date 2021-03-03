package leetcode

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	for j := 1; j <= k-1; j++ {
		for i := 0; i < len(nums)-j; i++ {
			if nums[i] < nums[i+1] {
				nums[i] = nums[i+1]
			}
		}
	}

	return nums[:len(nums)-k+1]
}

func maxSlidingWindow2(nums []int, k int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}

	biggest, index := nums[0], 0
	for i := 1; i < k; i++ {
		if nums[i] >= biggest {
			biggest = nums[i]
			index = i
		}
	}
	output := make([]int, len(nums)+1-k)
	output[0] = biggest

	for i := 1; i <= len(nums)-k; i++ {
		if nums[i+k-1] >= biggest {
			biggest = nums[i+k-1]
			index = i + k - 1
		} else if index < i {
			// recompute
			biggest = nums[i]
			index = i
			for j := i + 1; j < i+k; j++ {
				if nums[j] >= biggest {
					biggest = nums[j]
					index = j
				}
			}

		}

		output[i] = biggest
	}

	return output
}

func maxSlidingWindow3(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	output := make([]int, 0, len(nums)+1-k)
	deepQueue := make([]int, 0, k)

	for i := 0; i < len(nums); i++ {
		// try to insert data to queue
		{
			if len(deepQueue) == 0 || nums[i] <= nums[deepQueue[len(deepQueue)-1]] {
				deepQueue = append(deepQueue, i)
			}

			for j := 0; j < len(deepQueue); j++ {
				if nums[i] > nums[deepQueue[j]] {
					deepQueue[j] = i
					deepQueue = deepQueue[:j+1]
				}
			}
		}

		if i >= k-1 {
			output = append(output, nums[deepQueue[0]])

			if deepQueue[0] <= i-k+1 {
				deepQueue = deepQueue[1:]
			}
		}
	}

	return output
}

/*
输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
*/
