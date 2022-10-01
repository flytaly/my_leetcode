package longestpalindromicsubstring

func longestPalindrome(s string) string {
	result := s[0:1]

	// find longest palindrome starting from the center
	longest := func(pal string, l, r int) {
		for l >= 0 && r < len(s) {
			if s[l] != s[r] {
				return
			}
			pal = string(s[l]) + pal + string(s[r])
			if len(result) < len(pal) {
				result = pal
			}
			l, r = l-1, r+1
		}
	}

	for i := 1; i < len(s); i++ {
		longest(s[i:i+1], i-1, i+1) // odd length
		longest("", i-1, i)         // even length
	}

	return result
}
