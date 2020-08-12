package stack

func isValid(s string) bool {
	n := len(s)
	if n < 1 {
		return true
	}

	if n%2 != 0 {
		return false
	}

	dir := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := make([]byte, 0)
	for i := 0; i < n; i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
			continue
		}
		if dir[stack[len(stack)-1]] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}
