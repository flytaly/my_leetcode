package reducearraysizetothehalf

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestMaxSetSize(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}, 2},
		{[]int{7, 7, 7, 7, 7, 7}, 1},
		{[]int{1, 9}, 1},
		{[]int{9, 77, 63, 22, 92, 9, 14, 54, 8, 38, 18, 19, 38, 68, 58, 19}, 5},
	}

	for _, tt := range tests {
		got := minSetSize(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d: got %d, want %d", tt.nums, got, tt.want)
		}
	}
}

func BenchmarkMaxSetSize(b *testing.B) {
	n := 300
	nums := make([]int, n)
	rand.Seed(42)
	for i := 0; i < n; i++ {
		nums[i] = rand.Int()
	}
	for i := 0; i < b.N; i++ {
		minSetSize(nums)
	}
}
