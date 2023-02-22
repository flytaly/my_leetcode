package reducearraysizetothehalf

import (
	"reflect"
	"testing"
)

func TestMaxSetSize(t *testing.T) {
	tests := []struct {
		weights []int
		days    int
		want    int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 15},
		{[]int{3, 2, 2, 4, 1, 4}, 3, 6},
		{[]int{1, 2, 3, 1, 1}, 4, 3},
	}

	for _, tt := range tests {
		got := shipWithinDays(tt.weights, tt.days)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d (%d days): got %d, want %d", tt.weights, tt.days, got, tt.want)
		}
	}
}
