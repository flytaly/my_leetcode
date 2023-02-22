package longestconsecutivesequence

func longestConsecutive(nums []int) (result int) {
	numSet := map[int]struct{}{}

	for _, n := range nums {
		numSet[n] = struct{}{}
	}

	isNumExists := func(n int) bool {
		if _, exists := numSet[n]; exists {
			return true
		}
		return false
	}

	for _, n := range nums {
		if !isNumExists(n - 1) {
			streak := 1
			for currentNum := n + 1; isNumExists(currentNum); currentNum++ {
				streak += 1
			}
			if result < streak {
				result = streak
			}
		}
	}

	return result
}

// Get length of the sequence that starts with "start" number.
// if length isn't yet defined add it to the set of lengths.
func getLength(start int, lengths *map[int]int) int {
	if (*lengths)[start] > 0 {
		return (*lengths)[start]
	}

	if (*lengths)[start+1] == 0 { // there is no such number in the map => end of the sequence
		return 1
	}

	(*lengths)[start] = 1 + getLength(start+1, lengths)

	return (*lengths)[start]
}

func longestConsecutiveV1(nums []int) (result int) {
	lenMap := map[int]int{}

	for _, n := range nums {
		lenMap[n] = -1 // set length as temporary undefined
	}

	for _, n := range nums {
		l := getLength(n, &lenMap)
		if result < l {
			result = l
		}

	}

	return result
}
