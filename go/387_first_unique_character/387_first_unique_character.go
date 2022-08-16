package firstuniquecharacter

/*
	387. First Unique Character in a String

	Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.
*/

func firstUniqChar(s string) int {
	chars := map[rune]struct{}{}
	unique := map[rune]struct{}{}

	for _, v := range s {
		if _, has := chars[v]; has {
			delete(unique, v)
			continue
		}
		chars[v] = struct{}{}
		unique[v] = struct{}{}
	}

	for i, v := range s {
		if _, has := unique[v]; has {
			return i
		}
	}

	return -1
}
