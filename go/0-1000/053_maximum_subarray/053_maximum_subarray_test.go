package maximumsubarray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{-2, -5, 6, -2, -3, 1, 5, -6}, 7},
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
		{[]int{-2, 1}, 1},
	}

	for _, tt := range tests {
		got := maxSubArray(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d: got %d, want %d", tt.nums, got, tt.want)
		}
	}
	fmt.Println("err")
}
