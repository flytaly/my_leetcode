package palindromicsubstrings

func countSubstrings(s string) (result int) {
	count := func(l, r int) {
		for l >= 0 && r < len(s) {
			if s[l] != s[r] {
				return
			}
			result += 1
			l, r = l-1, r+1
		}
	}

	for i := range s {
		count(i-1, i+1) // odd length
		count(i-1, i)   // even length
	}

	return result + len(s)
}
