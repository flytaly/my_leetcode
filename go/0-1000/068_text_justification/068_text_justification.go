package textjustification

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	result := []string{}
	line := []string{}
	wordsLen := 0

	var pushLine = func(isLast bool) {
		var sb strings.Builder
		spaceLen := maxWidth - wordsLen
		spaceNum := len(line) - 1
		for i, w := range line {
			sb.WriteString(w)

			if isLast {
				if i < len(line)-1 {
					sb.WriteString(" ")
				}
				continue
			}

			if spaceNum > 0 {
				nextSpaceLen := (spaceLen + spaceNum - 1) / spaceNum // ceil
				sb.WriteString(strings.Repeat(" ", nextSpaceLen))
				spaceNum -= 1
				spaceLen -= nextSpaceLen
			}
		}

		if sb.Len() < maxWidth {
			sb.WriteString(strings.Repeat(" ", maxWidth-sb.Len()))
		}

		result = append(result, sb.String())
	}

	for i := 0; i < len(words); i++ {
		word := words[i]

		if wordsLen+len(word)+len(line)-1 < maxWidth {
			line = append(line, word)
			wordsLen += len(word)
			continue
		}

		pushLine(false)

		line = []string{}
		wordsLen = 0
		i -= 1
	}

	if len(line) > 0 {
		pushLine(true)
	}

	return result
}
