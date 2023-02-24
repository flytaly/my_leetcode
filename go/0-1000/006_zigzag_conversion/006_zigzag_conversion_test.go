package zigzag_conversion

import "testing"

func Test_convert(t *testing.T) {
	tests := []struct {
		s       string
		numRows int
		want    string
	}{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"A", 1, "A"},
		{"AB", 1, "AB"},
	}
	for _, tt := range tests {
		if got := convert(tt.s, tt.numRows); got != tt.want {
			t.Errorf("%v %v => %s, want %v", tt.s, tt.numRows, got, tt.want)
		}
	}
}
