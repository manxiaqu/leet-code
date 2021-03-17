package leetcode

// 递归超出时间限制
func maxArea(height []int) int {
	max := minInt(height[0], height[1])
	if len(height) == 2 {
		return max
	}

	biggest := make([]int, 1)
	biggest[0] = 0

	for i := 1; i < len(height); i++ {
		// cal area
		for j := 0; j < len(biggest); j++ {
			area := minInt(height[biggest[j]], height[i]) * (i - biggest[j])
			if area > max {
				max = area
			}
		}
		if height[i] > height[biggest[len(biggest)-1]] {
			biggest = append(biggest, i)
		}

	}

	return max
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 从头尾分别
func maxArea2(height []int) int {
	i, j, max := 0, len(height)-1, 0
	for i < j {
		heightI, heightJ := height[i], height[j]
		area := minInt(heightI, heightJ) * (j - i)
		if area > max {
			max = area
		}
		if heightI < heightJ {
			i++
		} else {
			j--
		}
	}

	return max
}
