/*
139. Word Break

Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.
*/

package wordbreak

func wordBreak(s string, wordDict []string) bool {

	dictMap := map[string]struct{}{}
	for _, w := range wordDict {
		dictMap[w] = struct{}{}
	}

	// subStrs is a slice of substrings, that indicates
	// if there is a way to break substring of s
	subStrs := make([]bool, len(s)+1)
	// just an inital value for the loop
	subStrs[0] = true

	for length, v := range subStrs {
		if v {
			accumulate := ""
			for i := length; i < len(s); i++ {
				accumulate += string(s[i])
				_, has := dictMap[accumulate]
				if has {
					subStrs[i+1] = true
				}
			}
		}

	}

	return subStrs[len(subStrs)-1]
}
