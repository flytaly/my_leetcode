package wordbreak

import (
	"reflect"
	"testing"
)

func TestWordBreak(t *testing.T) {
	tests := []struct {
		s        string
		wordDict []string
		want     bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
		{"aaaaaaa", []string{"aaaa", "aaa"}, true},
	}

	for _, tt := range tests {
		got := wordBreak(tt.s, tt.wordDict)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q => %q. got %t, want %t", tt.s, tt.wordDict, got, tt.want)
		}
	}
}
