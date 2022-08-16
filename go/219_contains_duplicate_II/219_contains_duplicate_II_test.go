package contains_duplicate_ii

import (
	"reflect"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
	}

	for _, tt := range tests {
		got := containsNearbyDuplicate(tt.nums, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("(%d, %d): got %t, want %t", tt.nums, tt.k, got, tt.want)
		}
	}
}
