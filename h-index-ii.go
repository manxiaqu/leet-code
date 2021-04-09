package leetcode

func hIndex(citations []int) int {
	n := 0

	for i := 0; i < len(citations); i++ {
		l := len(citations) - i
		if n >= l {
			break
		}

		if citations[i] < l {
			n = citations[i]
		} else {
			n = l
		}
	}

	return n
}
