package leetcode

func hammingWeight(num uint32) int {
	count := 0

	for num > 0 {
		if num%2 == 1 {
			count++
		}
		num = num / 2
	}

	return count
}

func hammingWeight2(num uint32) int {
	count := 0

	for num > 0 {
		num = num & (num - 1)
		count++
	}

	return count
}

func hammingWeight3(num uint32) int {
	num = (num & 0x55555555) + ((num >> 1) & 0x55555555)
	num = (num & 0x33333333) + ((num >> 2) & 0x33333333)
	num = (num & 0x0f0f0f0f) + ((num >> 4) & 0x0f0f0f0f)
	num = (num & 0x00ff00ff) + ((num >> 8) & 0x00ff00ff)
	num = (num & 0x0000ffff) + ((num >> 16) & 0x0000ffff)
	return int(num)
}
