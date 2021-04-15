package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[len(nums1)-m:], nums1[:m])

	i, j0, j1 := 0, len(nums1)-m, 0
	for ; j0 < len(nums1) && j1 < n; i++ {
		if nums1[j0] < nums2[j1] {
			nums1[i] = nums1[j0]
			j0++
		} else {
			nums1[i] = nums2[j1]
			j1++
		}
	}

	for j1 < n {
		nums1[i] = nums2[j1]
		j1++
		i++
	}
}
