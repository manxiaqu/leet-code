package leetcode

func isUnique(astr string) bool {
	exist := 0

	for _, v := range astr {
		check := 1 << (v - 'a')

		if exist&check == check {
			return false
		}

		exist = exist | check
	}

	return true
}
