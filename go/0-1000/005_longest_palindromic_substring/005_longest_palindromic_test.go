package longestpalindromicsubstring

import "testing"

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"example1", "babad", "bab"},
		{"example2", "cbbd", "bb"},
		{"example3", "bb", "bb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
