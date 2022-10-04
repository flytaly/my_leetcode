package palindromicsubstrings

import "testing"

func Test_countSubstrings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"example1", "abc", 3},
		{"example1", "aaa", 6}, // "a", "a", "a", "aa", "aa", "aaa"
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubstrings(tt.s); got != tt.want {
				t.Errorf("countSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
