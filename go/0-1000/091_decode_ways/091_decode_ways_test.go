package decodeways

import "testing"

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"12", 2},
		{"226", 3},
		{"1123", 5},
		{"123123", 9},
		{"06", 0},
		{"10", 1},
		{"2101", 1},
		{"207", 1},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := numDecodings(tt.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}
