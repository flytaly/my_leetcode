package validparentheses

import "testing"

func Test_isValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"example1", "()", true},
		{"example2", "()[]{}", true},
		{"example3", "(]", false},
		{"example4", "((2+3)*3+1)*2", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
