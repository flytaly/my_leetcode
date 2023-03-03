package lettercombination

func letterCombinations(digits string) []string {
	var combinations func(prev string, i int)
	var result = []string{}
	var buttons = map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	combinations = func(prev string, i int) {
		d := rune(digits[i])
		chars := buttons[d]
		for _, c := range chars {
			str := prev + string(c)
			if i >= len(digits)-1 {
				result = append(result, str)
				continue
			}
			combinations(str, i+1)
		}
	}

	if digits == "" {
		return result
	}

	combinations("", 0)

	return result
}
