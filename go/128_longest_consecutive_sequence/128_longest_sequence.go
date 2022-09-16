package longestconsecutivesequence

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

func longestConsecutive(nums []int) (result int) {
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
