package string

import (
	"math"
	"strconv"
	"strings"
)

func myAtoi(str string) int {
	n := len(str)
	if n < 1 {
		return 0
	}

	start := -1
	for i := 0; i < n; i++ {
		switch str[i] {
		case ' ', '+':
			continue
		case '-':
			start = i
			goto next
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			start = i
			goto next
		default:
			return 0
		}
	}

next:
	if start == -1 {
		return 0
	}

	if start > 0 && str[start] == '-' {
		if str[start-1] == '+' {
			return 0
		}
	}

	index := strings.Index(str, "+")
	if index < start {
		if index != -1 && index < n-1 {
			if str[index+1] == ' ' || str[index+1] == '+' {
				return 0
			}
		}
	}

	end := -1
	for i := start + 1; i < n; i++ {
		switch str[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			continue
		default:
			end = i
			goto finish
		}
	}

	if end == -1 {
		end = n - 1
		res := str[start : end+1]
		atoi, _ := strconv.Atoi(res)
		if atoi < math.MinInt32 {
			return math.MinInt32
		}
		if atoi > math.MaxInt32 {
			return math.MaxInt32
		}
		return atoi
	}

finish:
	res := str[start:end]
	atoi, _ := strconv.Atoi(res)
	if atoi < math.MinInt32 {
		return math.MinInt32
	}
	if atoi > math.MaxInt32 {
		return math.MaxInt32
	}
	return atoi
}

func myAtoi2(str string) int {
	n := len(str)
	if n < 1 {
		return 0
	}

	var symbol byte
	res := make([]byte, 0)
	for i := 0; i < n; i++ {
		if 'A' < str[i] && str[i] < 'z' {
			break
		}
		if str[i] == ' ' {
			if len(res) == 0 {
				continue
			} else {
				break
			}
		}
		if str[i] == '+' || str[i] == '-' {
			if len(res) == 0 {
				symbol = str[i]
				continue
			} else {
				break
			}
		}
		res = append(res, str[i])
	}

	if len(res) == 0 {
		return 0
	}

	ans := 0
	for i, j := len(res)-1, 0; i >= 0; i-- {
		ans += int(res[i]-'0') * int(math.Pow10(j))
		j++
	}

	if symbol == '+' {
		if ans > math.MaxInt32 {
			return math.MaxInt32
		}
		return ans
	} else {
		tmp := 0 - ans
		if tmp < math.MinInt32 {
			return math.MinInt32
		}
		return tmp
	}
}
