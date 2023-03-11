package lengthOfLastWord

func lengthOfLastWord(s string) (count int) {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != byte(' ') {
			count++
			continue
		}

		if count > 0 {
			break
		}
	}
	return count
}
