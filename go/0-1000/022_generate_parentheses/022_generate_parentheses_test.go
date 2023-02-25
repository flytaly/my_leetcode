package generateparentheses

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		n    int
		want []string
	}{
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{1, []string{"()"}},
	}

	for _, tt := range tests {
		got := generateParenthesis(tt.n)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v => got %v, want %v", tt.n, got, tt.want)
		}
	}
}
