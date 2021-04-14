package leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	// append 0 to first/end
	flowerbed = append(flowerbed, 0)
	flowerbed = append([]int{0}, flowerbed...)

	for i := 0; i < len(flowerbed)-2 && n != 0; {
		if flowerbed[i] == 0 && flowerbed[i+1] == 0 && flowerbed[i+2] == 0 {
			flowerbed[i+1] = 1
			n--
			i++
		}
		i++
	}

	return n == 0
}
