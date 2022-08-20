package min_stops

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		target    int
		startFuel int
		stations  [][]int
		want      int
	}{
		{1, 1, [][]int{}, 0},
		{100, 1, [][]int{{10, 100}}, -1},
		{100, 10, [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}, 2},
		{100, 50, [][]int{{25, 25}, {50, 50}}, 1},
	}

	for _, tt := range tests {
		got := minRefuelStops(tt.target, tt.startFuel, tt.stations)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d, %d, %d: got %d, want %d", tt.target, tt.startFuel, tt.stations, got, tt.want)
		}
	}
}
