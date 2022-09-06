package countingbits

import (
	"reflect"
	"testing"
)

func Test_countBits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		// {"e1", 2, []int{0, 1, 1}},
		// {"e2", 5, []int{0, 1, 1, 2, 1, 2}},
		{"e2", 8, []int{0, 1, 1, 2, 1, 2, 2, 3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBits(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
