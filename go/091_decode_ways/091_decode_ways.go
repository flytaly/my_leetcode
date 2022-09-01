package decodeways

import "strconv"

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	sum1 := 1
	sum2 := 1

	for i := 1; i < len(s); i++ {
		sum := 0

		if s[i] != '0' {
			sum = sum1
		}

		pair, _ := strconv.Atoi(s[i-1 : i+1])
		if pair >= 10 && pair <= 26 {
			sum += sum2
		}

		sum2 = sum1
		sum1 = sum
	}

	return sum1
}
