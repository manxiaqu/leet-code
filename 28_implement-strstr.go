package leetcode

import "bytes"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	hbytes := []byte(haystack)
	nbytes := []byte(needle)

	for i := 0; i <= len(hbytes)-len(nbytes); i++ {
		if bytes.Equal(hbytes[i:i+len(nbytes)], nbytes[:]) {
			return i
		}
	}

	return -1
}
