package longestsubstringwithoutrepeat

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) (result int) {
	letters := map[rune]int{} // map letter to its latest index
	start := 0
	for i, ch := range s {
		if letters[ch] > start {
			start = letters[ch]
		}
		result = max(result, i+1-start)
		letters[ch] = i + 1
	}

	return result
}

func lengthOfLongestSubstringV1(s string) (longest int) {
	visited := map[rune]int{}

	for i, char := range s {
		if startIndex, ok := visited[char]; ok {
			for k, v := range visited {
				if v <= startIndex {
					delete(visited, k)
				}
			}
		}

		visited[char] = i
		if longest < len(visited) {
			longest = len(visited)
		}
	}

	return longest
}
