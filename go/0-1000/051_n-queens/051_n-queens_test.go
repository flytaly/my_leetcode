package nqueens

import (
	"reflect"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]string
	}{
		{"1x1", 1, [][]string{{"Q"}}},
		{"2x2", 2, nil},
		{"3x3", 3, nil},
		{"4x4", 4, [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}},
		{"6x6", 6, [][]string{{".Q....", "...Q..", ".....Q", "Q.....", "..Q...", "....Q."}, {"..Q...", ".....Q", ".Q....", "....Q.", "Q.....", "...Q.."}, {"...Q..", "Q.....", "....Q.", ".Q....", ".....Q", "..Q..."}, {"....Q.", "..Q...", "Q.....", ".....Q", "...Q..", ".Q...."}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveNQueens(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
