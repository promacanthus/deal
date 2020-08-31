package string

func reverseString(s []byte) {
	n := len(s)
	if n < 2 {
		return
	}
	i, j := 0, n-1
	for i < j {
		s[i], s[n-i-1] = s[n-i-1], s[i]
		i++
		j--
	}
}
