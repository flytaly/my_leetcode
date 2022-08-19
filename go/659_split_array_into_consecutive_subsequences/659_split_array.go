package splitArrayIntoConsecutiveSubsequences

/*
	659. Split Array into Consecutive Subsequences
	https://leetcode.com/problems/split-array-into-consecutive-subsequences/
*/

func isPossible(nums []int) bool {
	freq := map[int]int{} // frequencies
	for _, v := range nums {
		freq[v]++
	}

	tails := map[int]int{} // amount of subsequences that could concatenate the number

	for _, start := range nums {
		if freq[start] <= 0 {
			continue
		}
		freq[start]--
		if tails[start] > 0 { // join number as a tail to some previous subsequence
			tails[start] -= 1
			tails[start+1] += 1
			continue
		}
		end := start + 2
		for i := start + 1; i <= end; i++ { // take 3 subsequent numbers if possible
			if freq[i] == 0 {
				return false
			}
			freq[i] -= 1
		}
		tails[end+1] += 1
	}

	return true
}
