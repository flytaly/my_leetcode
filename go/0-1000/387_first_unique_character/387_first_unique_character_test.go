package firstuniquecharacter

import (
	"reflect"
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"leetcode", 0},
		{"loveleetcode", 2},
		{"aabb", -1},
	}

	for _, tt := range tests {
		got := firstUniqChar(tt.s)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s: got %d, want %d", tt.s, got, tt.want)
		}
	}
}
