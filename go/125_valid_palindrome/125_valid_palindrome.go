package validpalindrome

import "strings"

func isValidChar(c int) bool {
	if c >= int('a') && c <= int('z') {
		return true
	}
	return c >= int('0') && c <= int('9')
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	for left, right := 0, len(s)-1; left < right; {
		if !isValidChar(int(s[left])) {
			left++
			continue
		}
		if !isValidChar(int(s[right])) {
			right--
			continue
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}
