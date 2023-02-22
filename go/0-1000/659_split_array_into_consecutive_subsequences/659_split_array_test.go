package splitArrayIntoConsecutiveSubsequences

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 3, 3, 4, 5}, true},
		{[]int{1, 2, 3, 3, 4, 4, 5, 5}, true},
		{[]int{1, 2, 3, 4, 4, 5}, false},
		{[]int{1, 2, 3, 5, 6, 7}, true},
	}

	for _, tt := range tests {
		got := isPossible(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d: got %t, want %t", tt.nums, got, tt.want)
		}
	}
}
