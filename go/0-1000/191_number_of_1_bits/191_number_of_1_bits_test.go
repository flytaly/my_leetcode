package numberof1bits

import "testing"

func Test_hammingWeight(t *testing.T) {
	tests := []struct {
		name string
		num  uint32
		want int
	}{
		{"e1", 0b1011, 3},
		{"e2", 0b10000000, 1},
		{"e3", 0b11111111111111111111111111111101, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.num); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
