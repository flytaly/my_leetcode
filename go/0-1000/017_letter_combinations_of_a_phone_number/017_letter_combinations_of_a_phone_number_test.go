package lettercombination

import (
	"reflect"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	tests := []struct {
		digits string
		want   []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"", []string{}},
	}
	for _, tt := range tests {
		if got := letterCombinations(tt.digits); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("letterCombinations() = %v, want %v", got, tt.want)
		}
	}
}
