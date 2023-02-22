package validparentheses

var opening = map[rune]rune{'[': ']', '(': ')', '{': '}'}
var closing = map[rune]rune{']': '[', ')': '(', '}': '{'}

func isValid(s string) bool {
	stack := []rune{}
	for _, c := range s {
		if opening[c] != 0 {
			stack = append(stack, c)
		}
		if closing[c] != 0 {
			l := len(stack)
			if l == 0 || stack[l-1] != closing[c] {
				return false
			}
			stack = stack[:l-1]
		}
	}

	return len(stack) == 0
}
