package generateparentheses

func generateParenthesis(n int) []string {
	result := []string{}

	var generate func(s string, open, close int)

	generate = func(s string, open, close int) {
		if close >= n {
			result = append(result, s)
		}
		if open < n {
			generate(s+"(", open+1, close)
		}
		if open > close {
			generate(s+")", open, close+1)
		}
	}

	generate("", 0, 0)

	return result
}
