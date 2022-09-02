package setmatrixzeroes

import (
	"reflect"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{"3x3", [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{"3x4", [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
		{
			"3x5",
			[][]int{{-4, -2147483648, 6, -7, 0}, {-8, 6, -8, -6, 0}, {2147483647, 2, -9, -6, -10}},
			[][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {2147483647, 2, -9, -6, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if setZeroes(tt.matrix); !reflect.DeepEqual(tt.matrix, tt.want) {
				t.Errorf("setZeroes() = %v, want %v", tt.matrix, tt.want)
			}
		})
	}
}
