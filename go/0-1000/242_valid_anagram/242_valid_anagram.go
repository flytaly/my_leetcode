package validanagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letters := map[rune]int{}

	for _, c := range s {
		letters[c]++
	}

	for _, c := range t {
		if letters[c] <= 0 {
			return false
		}
		letters[c]--
	}

	return true
}
