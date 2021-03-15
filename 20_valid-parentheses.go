package leetcode

func isValid(s string) bool {

	stringBytes := []byte(s)

	stack := make([]byte, 0)

	for _, b := range stringBytes {
		if b == ')' || b == ']' || b == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != shouldBe(b) {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, b)
		}
	}

	return len(stack) == 0
}

func shouldBe(b byte) byte {
	if b == ')' {
		return '('
	}
	if b == '}' {
		return '{'
	}
	if b == ']' {
		return '['
	}

	return byte(0)
}
