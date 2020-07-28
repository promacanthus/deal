package DP

func isMatch(s string, p string) bool {
	sLength := len(s)
	pLength := len(p)
	if sLength < 1 || pLength < 1 {
		return false
	}

	sBytes := []byte(s)
	pBytes := []byte(p)

	for i := 0; i < pLength; i++ {
		if pBytes[i] >= 'a' && pBytes[i] <= 'z' {
			if sBytes[i] == pBytes[i] {
				continue
			}
		} else {
			switch pBytes[i] {
			case '.':
				continue
			case '*':
				if i == 0 {
					continue
				}
			}
		}

	}
	return true
}
