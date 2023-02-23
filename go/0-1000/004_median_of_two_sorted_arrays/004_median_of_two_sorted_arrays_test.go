package medianoftwosortedarrays

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5, 6, 7, 8}, 4},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4, 5, 6, 7, 8}, 3.5},
		{[]int{2}, []int{1}, 1.5},
		{[]int{2}, []int{1, 3}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{}, []int{2, 2}, 2},
		{[]int{}, []int{1}, 1},
	}
	for _, tt := range tests {
		if got := findMedianSortedArrays(tt.nums1, tt.nums2); got != tt.want {
			t.Errorf("%v & %v => %v, want %v", tt.nums1, tt.nums2, got, tt.want)
		}
	}
}
