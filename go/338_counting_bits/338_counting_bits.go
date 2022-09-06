package countingbits

// See 191th problem "Number of 1 bits"
func countBits(n int) []int {
	res := make([]int, n+1)

	for i := 1; i < len(res); i++ {
		res[i] = res[i&(i-1)] + 1
	}

	return res
}
