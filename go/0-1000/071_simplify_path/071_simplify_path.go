package simplifypath

import "strings"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func simplifyPath(path string) string {
	split := strings.Split(path, "/")
	dirs := make([]string, len(split))
	index := 0
	for _, dir := range split {
		switch dir {
		case ".":
		case "":
		case "..":
			index = max(0, index-1)
		default:
			index += 1
			dirs[index] = dir
		}
	}

	if index == 0 {
		return "/"
	}

	res := strings.Join(dirs[:index+1], "/")
	return res
}
