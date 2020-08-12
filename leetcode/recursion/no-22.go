package recursion

// backtracking
func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	result := make([]string, 0)
	gp("(", 1, n, 1, 0, &result)
	return result
}

func gp(cur string, i, n, open, close int, result *[]string) {
	if i == n*2 {
		*result = append(*result, cur)
		return
	}

	if open < n {
		cur += "("
		gp(cur, i+1, n, open+1, close, result)
		cur = cur[:len(cur)-1]
	}

	if close < open {
		cur += ")"
		gp(cur, i+1, n, open, close+1, result)
		cur = cur[:len(cur)-1]
	}
}

