package firstmissingpositive

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 1, 0, 2}, 4},
		{[]int{3, 4, -1, 1}, 2},
		{[]int{7, 8, 9, 11, 12}, 1},
	}
	for _, tt := range tests {
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)
		if got := firstMissingPositive(nums); got != tt.want {
			t.Errorf("firstMissingPositive(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
