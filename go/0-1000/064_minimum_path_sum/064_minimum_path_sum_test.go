package minimumpathsum

import "testing"

func Test_minPathSum(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			"ex1",
			[][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			7,
		},
		{
			"ex2",
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum(tt.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
