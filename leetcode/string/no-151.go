package string

import (
	"strings"
)

func reverseWords(s string) string {
	n := len(s)
	if n < 1 {
		return ""
	}

	ans := make([]string, 0)
	start, end := -1, -1
	for i := 0; i < n; i++ {
		if s[i] != ' ' && start == -1 {
			start = i
		}
		if s[i] == ' ' && start != -1 && end == -1 {
			end = i - 1
			ans = append(ans, s[start:end+1])
			start = -1
			end = -1
		}
		if i == n-1 && start != -1 && end == -1 {
			end = i
			ans = append(ans, s[start:end+1])
		}
	}

	if len(ans) == 0 {
		return ""
	}

	res := ""
	for i := len(ans) - 1; i >= 0; i-- {
		res += ans[i]
		res += " "
	}
	return res[:len(res)-1]
}

func reverseWordsBuiltIn(s string) string {
	n := len(s)
	if n < 1 {
		return ""
	}

	splits := strings.Split(s, " ")

	if len(splits) == 0 {
		return ""
	}

	res := ""
	for i := len(splits) - 1; i >= 0; i-- {
		if splits[i] != "" {
			res += splits[i]
			res += " "
		}

	}

	if res == "" {
		return res
	}

	return res[:len(res)-1]
}

func reverseWordss(s string) string {
	n := len(s)
	if n < 1 {
		return ""
	}

	res := make([]string, 0)
	word := make([]byte, 0)
	for i := 0; i < n; i++ {
		if s[i] == ' ' {
			if len(word) == 0 {
				continue
			}
			res = append(res, string(word))
			word = make([]byte, 0)
		} else {
			word = append(word, s[i])
		}
	}

	if len(word) != 0 {
		res = append(res, string(word))
	}
	if len(res) == 0 {
		return ""
	}

	ans := ""
	for i := len(res) - 1; i >= 0; i-- {
		ans += res[i]
		ans += " "
	}
	return ans[:len(ans)-1]
}
