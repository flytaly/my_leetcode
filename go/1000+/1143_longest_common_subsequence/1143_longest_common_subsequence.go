package longestcommonsubsequence

/*
	1143. Longest Common Subsequence
*/

/*
            ┌──────[a0,b0]──────┐
        [a0, b1]            [a1, b0]
       ┌────┴────┐         ┌────┴────┐
   [a0, b2]  [a1, b1]===[a1, b1] [a2, b0]
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var memo [][]int

// recursion with memoization
func lcsMemo(s1, s2 string, i, j int) int {
	if i >= len(s1) || j >= len(s2) {
		return 0
	}

	if s1[i] == s2[j] {
		m := 1 + lcsMemo(s1, s2, i+1, j+1)
		memo[i][j] = m
		return m
	}

	l := memo[i+1][j]
	r := memo[i][j+1]

	if l == -1 {
		l = lcsMemo(s1, s2, i+1, j)
	}
	if r == -1 {
		r = lcsMemo(s1, s2, i, j+1)

	}

	m := max(l, r)

	memo[i][j] = m

	return m
}

// DP
func lcsDP(s1, s2 string) int {
	memo := make([][]int, len(s1)+1)
	for i := range memo {
		memo[i] = make([]int, len(s2)+1)
	}

	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				memo[i+1][j+1] = 1 + memo[i][j]
				continue
			}
			memo[i+1][j+1] = max(memo[i][j+1], memo[i+1][j])
		}
	}

	return memo[len(s1)][len(s2)]
}

func longestCommonSubsequence(text1 string, text2 string) int {

	/* memo = make([][]int, len(text1)+1)
	for i := range memo {
		memo[i] = make([]int, len(text2)+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return lcsMemo(text1, text2, 0, 0) */

	return lcsDP(text1, text2)
}
