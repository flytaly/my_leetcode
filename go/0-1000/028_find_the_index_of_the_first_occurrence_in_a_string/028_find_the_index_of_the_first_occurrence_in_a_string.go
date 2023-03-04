package findsubstr

// import (
//     "strings"
// )
//
// func strStr(haystack string, needle string) int {
//     return strings.Index(haystack, needle)
//}

func strStr(haystack string, needle string) int {
	poss := make(map[int]struct{})

	for i := 0; i < len(haystack); i++ {
		s := haystack[i]
		poss[i] = struct{}{}
		for startIdx := range poss {
			idx := i - startIdx
			if needle[idx] != s {
				delete(poss, startIdx)
				continue
			}
			if idx == len(needle)-1 {
				return startIdx
			}
		}
	}

	return -1
}
