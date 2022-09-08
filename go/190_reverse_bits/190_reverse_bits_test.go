package reversebits

import "testing"

func Test_reverseBits(t *testing.T) {
	tests := []struct {
		name string
		num  uint32
		want uint32
	}{
		{"ex1", 0b00000010100101000001111010011100, 0b00111001011110000010100101000000},
		{"ex2", 0b11111111111111111111111111111101, 0b10111111111111111111111111111111},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits(tt.num); got != tt.want {
				t.Errorf("reverseBits() = %b, want %b", got, tt.want)
			}
		})
	}
}
