package hashtable

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n < 1 {
		return n
	}
	m := make(map[byte]int)
	ans := 0
	for i, j := 0, 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for j < n && m[s[j]] == 0 {
			m[s[j]]++
			j++
		}
		ans = max(ans, j-i)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
