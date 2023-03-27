package spiralmatrixII

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]int
	}{
		{"example_1", 3, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
		{"example_2", 1, [][]int{{1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMatrix(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("(%v) => %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
