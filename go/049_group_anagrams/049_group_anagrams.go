package groupanagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	groups := map[string][]string{} // map: sorted str => original values

	for _, str := range strs {
		slice := []rune(str)
		sort.Slice(slice, func(i int, j int) bool { return slice[i] < slice[j] })
		s := string(slice)
		groups[s] = append(groups[s], str)
	}

	result := make([][]string, len(groups))

	i := 0
	for _, g := range groups {
		result[i] = g
		i++
	}

	return result
}
