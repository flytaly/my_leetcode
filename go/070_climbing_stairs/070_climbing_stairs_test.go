package climbingstairs

import (
	"reflect"
	"testing"
)

func TestClib(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{2, 2},
		{3, 3},
		{4, 5},
	}

	for _, tt := range tests {
		got := climbStairs(tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d: got %d, want %d", tt.n, got, tt.want)
		}
	}
}
