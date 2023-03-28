package uniquepathsII

import "testing"

func Test_uniquePathsWithObstacles(t *testing.T) {
	tests := []struct {
		name         string
		obstacleGrid [][]int
		want         int
	}{
		{"ex1", [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2},
		{"ex2", [][]int{{0, 1}, {0, 0}}, 1},
		{"ex3", [][]int{{0, 0, 0}, {0, 1, 0}, {0, 1, 0}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles(%v) = %v, want %v", tt.obstacleGrid, got, tt.want)
			}
		})
	}
}
