package leetcode

func trap(height []int) int {
	// index
	left, right, value, pos := 0, len(height)-1, 0, 1

	for left < right {
		// process the small side
		if height[left] < height[right] {
			// check pos and left
			pos = left + 1
			if height[pos] < height[left] {
				value = value + height[left] - height[pos]
				height[pos] = height[left]
			}
			left = pos
		} else {
			pos = right - 1
			if height[pos] < height[right] {
				value = value + height[right] - height[pos]
				height[pos] = height[right]
			}
			right = pos
		}
	}

	return value
}
