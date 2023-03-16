package plusone

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	tests := []struct {
		digits []int
		want   []int
	}{
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{9}, []int{1, 0}},
		{[]int{1, 9, 9}, []int{2, 0, 0}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
	}
	for _, tt := range tests {
		if got := plusOne(tt.digits); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("plusOne(%v) = %v, want %v", tt.digits, got, tt.want)
		}
	}
}
