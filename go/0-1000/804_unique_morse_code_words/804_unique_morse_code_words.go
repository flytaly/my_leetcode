package uniquemorsecodewords

/* 804. Unique Morse Code Words */

var morse = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func encode(word string) (code string) {
	for i := 0; i < len(word); i++ {
		char := word[i]
		code += morse[char-byte('a')]
	}
	return code
}

func uniqueMorseRepresentations(words []string) int {
	codes := map[string]struct{}{}

	for _, w := range words {
		codes[encode(w)] = struct{}{}
	}

	return len(codes)
}
