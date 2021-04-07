package leetcode

func search(nums []int, target int) bool {
	if nums[0] == target {
		return true
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] && (nums[0] < target || target < nums[i]) {
			return false
		}

		if target == nums[i] {
			return true
		}
	}

	return false
}

func search2(nums []int, target int) bool {
	if nums[0] == target {
		return true
	}

	// search from start or search from end
	if nums[0] > target {
		for i := len(nums) - 1; i > 0 && target >= nums[i]; i-- {
			if target == nums[i] {
				return true
			}
		}

		return false
	}

	for i := 1; i < len(nums) && target <= nums[i]; i++ {
		if target == nums[i] {
			return true
		}
	}

	return false
}
