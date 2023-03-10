package removeduplicates

import "testing"

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 2}, 2},                      // => {1, 2, _}
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5}, // => {0,1,2,3,4,_,_,_,_,_}
	}
	for _, tt := range tests {
		if got := removeDuplicates(tt.nums); got != tt.want {
			t.Errorf("removeDuplicates(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
