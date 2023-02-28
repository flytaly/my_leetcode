package reverseinteger

import "testing"

func Test_reverse(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{1534236469, 0},
		{-2147483412, -2143847412},
	}
	for _, tt := range tests {
		if got := reverse(tt.x); got != tt.want {
			t.Errorf("reverse() = %v, want %v", got, tt.want)
		}
	}
}
