package longestsubstringwithoutrepeat

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"ex1", "abcabcbb", 3},
		{"ex2", "bbbbb", 1},
		{"ex3", "pwwkew", 3},
		{"ex4", "aab", 2},
		{"ex5", "dvdf", 3},
		{"ex6", "d", 1},
		{"ex7", "", 0},
		{"ex8", "au", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
