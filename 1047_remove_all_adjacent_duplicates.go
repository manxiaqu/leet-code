package leetcode

func removeDuplicates(S string) string {
	// 向处理栈一样的方式处理字符串
	output := ""

	for i := 0; i < len(S); i++ {
		if output == "" || output[len(output)-1] != S[i] {
			output = output + string(S[i])
			continue
		}

		output = output[:len(output)-1]
	}

	return output
}

func removeDuplicates2(S string) string {
	if len(S) < 2 {
		return S
	}
	output, ori := make([]byte, 0, len(S)), []byte(S)
	for i := 0; i < len(ori); i++ {
		if len(output) == 0 || output[len(output)-1] != ori[i] {
			output = append(output, ori[i])
			continue
		}

		output = output[:len(output)-1]
	}

	return string(output)
}
