package leetcode

import "fmt"

func largestNumber(nums []int) string {
	// i is not ordered, j is ordered
	for i, j := 1, 0; i < len(nums); i++ {
		for k := 0; k <= j; k++ {
			// i should insert to k
			if bigger(nums[i], nums[k]) {
				t := nums[i]
				for c := j; c >= k; c-- {
					nums[c+1] = nums[c]
				}
				nums[k] = t
				break
			}
		}
		j++
	}
	if nums[0] == 0 {
		return "0"
	}

	output := ""

	for i := 0; i < len(nums); i++ {
		output = output + fmt.Sprintf("%d", nums[i])
	}

	return output
}

// returns if ab > ba
func bigger(a, b int) bool {
	if a == b {
		return true
	}

	as := []byte(fmt.Sprintf("%d%d", a, b))
	bs := []byte(fmt.Sprintf("%d%d", b, a))

	for i := 0; i < len(as); i++ {
		if as[i] > bs[i] {
			return true
		}
		if as[i] < bs[i] {
			return false
		}
	}

	return true
}
