package palindromenumber

import (
	"reflect"
	"testing"
)

func TestPalindrom(t *testing.T) {
	tests := []struct {
		x    int
		want bool
	}{
		{1, true},
		{121, true},
		{-121, false},
		{10, false},
	}

	for _, tt := range tests {
		got := isPalindrome(tt.x)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d -> %v, want %v", tt.x, got, tt.want)
		}
	}
}
