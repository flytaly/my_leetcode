package uniquemorsecodewords

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		words []string
		want  int
	}{
		{[]string{"gin", "zen", "gig", "msg"}, 2},
		{[]string{"a"}, 1},
	}

	for _, tt := range tests {
		got := uniqueMorseRepresentations(tt.words)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s: got %d, want %d", tt.words, got, tt.want)
		}
	}
}
