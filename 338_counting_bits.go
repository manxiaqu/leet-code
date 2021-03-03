package leetcode

func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	}
	output := make([]int, num+1)

	output[0] = 0
	output[1] = 1

	modBase := 2

	for i := 2; i <= num; i++ {
		index := i % modBase

		output[i] = 1 + output[index]

		if index == modBase-1 {
			modBase *= 2
		}
	}

	return output
}

func countBits2(num int) []int {
	if num == 0 {
		return []int{0}
	}
	output := make([]int, num+1)

	output[0] = 0

	for i := 1; i <= num; i++ {
		output[i] = 1 + output[i&(i-1)]
	}

	return output
}

/*
0 => 0
1 => 1
2 => 1
3 => 2
4 => 1
5 => 2
*/
