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

func lcsMemo(s1, s2 string, i1, i2 int) int {
	if i1 >= len(s1) || i2 >= len(s2) {
		return 0
	}

	if s1[i1] == s2[i2] {
		m := 1 + lcsMemo(s1, s2, i1+1, i2+1)
		memo[i1][i2] = m
		return m
	}

	l := memo[i1+1][i2]
	r := memo[i1][i2+1]

	if l == -1 {
		l = lcsMemo(s1, s2, i1+1, i2)
	}
	if r == -1 {
		r = lcsMemo(s1, s2, i1, i2+1)

	}

	m := max(l, r)

	memo[i1][i2] = m

	return m
}

func longestCommonSubsequence(text1 string, text2 string) int {

	memo = make([][]int, len(text1)+1)
	for i := range memo {
		memo[i] = make([]int, len(text2)+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	res := lcsMemo(text1, text2, 0, 0)

	return res
}
