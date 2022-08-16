package containsduplicate

import (
	"reflect"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, tt := range tests {
		got := containsDuplicate(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d: got %t, want %t", tt.nums, got, tt.want)
		}
	}
}
