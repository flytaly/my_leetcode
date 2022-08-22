package verifyinganaliendictionary

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isAlienSorted(words []string, order string) bool {
	charMap := map[byte]int{}

	for i := range order {
		charMap[order[i]] = i
	}

	word1 := words[0]
	for _, word2 := range words[1:] {
		isEqual := true
		for i := 0; i < min(len(word1), len(word2)); i++ {
			lIndex := charMap[word1[i]]
			rIndex := charMap[word2[i]]
			if lIndex > rIndex {
				return false
			}
			if lIndex < rIndex {
				isEqual = false
				break
			}
		}

		if isEqual && len(word1) > len(word2) {
			return false
		}

		word1 = word2
	}

	return true
}
