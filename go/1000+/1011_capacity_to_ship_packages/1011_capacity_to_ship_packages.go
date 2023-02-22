package reducearraysizetothehalf

func shipWithinDays(weights []int, days int) int {
	var minCap, maxCap int
	for _, v := range weights {
		maxCap += v
		if v > minCap {
			minCap = v
		}

	}

	// Binary search: check if middle point (MaxCap+MinCap)/2 has enough capacity.
	// If true -> use it as a new MaxCap, otherwise as a new MinCap

	var isEnough = func(capacity int) bool {
		sum := 0
		dayNum := 1
		for _, v := range weights {
			sum += v
			if sum > capacity {
				sum = v
				dayNum++
			}
			if dayNum > days {
				return false
			}
		}

		return true
	}

	var middle int = 0

	for {
		middle = (minCap + maxCap) / 2
		if middle == maxCap || middle == minCap {
			break
		}
		if isEnough(middle) {
			maxCap = middle
		} else {
			minCap = middle
		}
	}

	if isEnough(minCap) {
		return minCap
	}

	return maxCap
}
