package minimumwindowsubstring

func minWindow(s string, t string) string {
	minStart, minEnd := 0, 0 // indeces of the resulting minimal interval
	start, end := 0, 0       // sliding window pointers

	lettersLeft := map[byte]int{}
	for _, l := range t {
		lettersLeft[byte(l)] += 1
	}

	remain := len(t)            // indicates how many letters from "t" remain unused
	excessive := map[byte]int{} // store repeating letters

	for end < len(s) {
		last := s[end]
		end += 1

		if _, ok := lettersLeft[last]; !ok {
			continue
		}

		if lettersLeft[last] > 0 {
			lettersLeft[last] -= 1
			remain -= 1
		} else {
			excessive[last] += 1
		}

		for remain == 0 && start < end { // move left side of the window
			prev := minEnd - minStart
			if prev == 0 || (end-start) < prev { // save non-zero minimal interval
				minStart = start
				minEnd = end
			}
			first := s[start]
			start += 1
			if _, ok := lettersLeft[first]; !ok {
				continue
			}
			if excessive[first] > 0 { // there are "excessive" letters left => shrink window
				excessive[first] -= 1
				continue
			}
			lettersLeft[first] += 1
			remain += 1
		}

		if start >= end {
			break
		}
	}

	return s[minStart:minEnd]
}
