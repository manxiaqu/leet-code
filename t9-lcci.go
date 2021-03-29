package leetcode

func getValidT9Words(num string, words []string) []string {
	// strMap := [8][]byte{
	// 	// 2
	// 	[]byte{'a', 'b', 'c'},
	// 	// 3
	// 	[]byte{'d', 'e', 'f'},
	// 	// 4
	// 	[]byte{'g', 'h', 'i'},
	// 	// 5
	// 	[]byte{'j', 'k', 'l'},
	// 	// 6
	// 	[]byte{'m', 'n', 'o'},
	// 	// 7
	// }
	output := make([]string, 0)

	for _, word := range words {
		if len(num) != len(word) {
			continue
		}

		// check word char with num
		nBytes := []byte(num)
		// wBytes := []byte(word)
		add := true
		for i := 0; i < len(nBytes); i++ {
			if nBytes[i] <= '1' || nBytes[i] > '9' {
				add = false
				break
			}

			index := nBytes[i] - '2'
			start := index * 3
			end := start + 2
			if index == 5 {
				end = end + 1
			} else if index == 6 {
				start = start + 1
				end = end + 1
			} else if index == 7 {
				start = start + 1
				end = end + 2
			}

			if start > word[i]-'a' || word[i]-'a' > end {
				add = false
				break
			}
		}

		if add {
			output = append(output, word)
		}
	}

	return output
}

func getValidT9Words2(num string, words []string) []string {
	strMap := [8][]byte{
		// 2
		{0, 2},
		// 3
		{3, 5},
		// 4
		{6, 8},
		// 5
		{9, 11},
		// 6
		{11, 14},
		// 7
		{15, 18},
		// 8
		{19, 21},
		// 9
		{22, 25},
	}
	output := make([]string, 0)

	for _, word := range words {
		if len(num) != len(word) {
			continue
		}

		// check word char with num
		add := true
		for i, b := range []byte(num) {
			if b <= '1' || b > '9' {
				add = false
				break
			}

			if word[i]-'a' < strMap[b-'2'][0] || word[i]-'a' > strMap[b-'2'][1] {
				add = false
				break
			}
		}

		if add {
			output = append(output, word)
		}
	}

	return output
}
