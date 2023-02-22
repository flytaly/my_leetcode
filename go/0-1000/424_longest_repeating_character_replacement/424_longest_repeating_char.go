package longestrepeatingcharacterreplacement

// returns the most frequent element in the slice
func mostFreq(freq *[]int) (result int) {
	for _, v := range *freq {
		if v > result {
			result = v
		}
	}
	return result
}

func characterReplacement(s string, k int) (result int) {
	start := 0              // starting index of the Sliding Window
	windowMax := 0          // the maximum repetition in the Window
	freq := make([]int, 26) // map a letter to its frequency

	for end := range s { // end of the Window
		l := end - start + 1 // Window's length
		char := s[end] - 'A'

		freq[char] += 1
		if windowMax < freq[char] {
			windowMax = freq[char]
		}

		for l-windowMax > k { // move left side of the Window
			freq[s[start]-'A'] -= 1
			start += 1
			l -= 1
			windowMax = mostFreq(&freq) // can be optimized probably to not loop every time?
		}

		if result < l {
			result = l
		}
	}

	return result
}
