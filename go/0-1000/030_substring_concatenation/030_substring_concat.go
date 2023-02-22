package substringWithConcatenation

/*
30. Substring with Concatenation of All Words

You are given a string s and an array of strings words of the same length.
Return all starting indices of substring(s) in s that is a concatenation of each
word in words exactly once, in any order, and without any intervening characters.
*/

func findSubstring(s string, words []string) []int {
	result := []int{}
	wordSize := len(words[0])
	wordsNum := len(words)

	wordMap := map[string]int{}
	for _, v := range words {
		wordMap[v] = wordMap[v] + 1
	}

	for shift := 0; shift <= len(s)-wordSize*wordsNum; shift++ {
		// check if the prefix of the substring contains every words
		usedWordMap := map[string]int{}
		wordCount := 0
		for i := shift; i <= len(s)-wordSize; i += wordSize {
			word := s[i : i+wordSize]
			total := wordMap[word]
			used := usedWordMap[word]
			if used >= total {
				break
			}
			usedWordMap[word] += 1
			wordCount += 1
			if wordCount == wordsNum {
				result = append(result, shift)
			}
		}
	}

	return result
}
