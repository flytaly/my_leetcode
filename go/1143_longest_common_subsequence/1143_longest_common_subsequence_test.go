package longestcommonsubsequence

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		text1 string
		text2 string
		want  int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
		{"bsbininm", "jmjkbkjkv", 1},
	}

	for _, tt := range tests {
		got := longestCommonSubsequence(tt.text1, tt.text2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("(%s, %s) -> %d, want %d", tt.text1, tt.text2, got, tt.want)
		}
	}
}
