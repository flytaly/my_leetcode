package verifyinganaliendictionary

import (
	"reflect"
	"testing"
)

func TestIsAlienSorted(t *testing.T) {
	tests := []struct {
		words []string
		order string
		want  bool
	}{
		{[]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz", true},
		{[]string{"zirqhpfscx", "zrmvtxgelh", "vokopzrtc", "nugfyso", "rzdmvyf", "vhvqzkfqis", "dvbkppw", "ttfwryy", "dodpbbkp", "akycwwcdog"}, "khjzlicrmunogwbpqdetasyfvx", false},
	}

	for _, tt := range tests {
		got := isAlienSorted(tt.words, tt.order)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s->%s: got %t, want %t", tt.words, tt.order, got, tt.want)
		}
	}
}
