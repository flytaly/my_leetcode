package substringWithConcatenation

import (
	"reflect"
	"testing"
)

func TestFindSubstring(t *testing.T) {

	tests := []struct {
		s     string
		words []string
		want  []int
	}{
		{"barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
		{"barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}, []int{}},
		{"barfoofoobarthefoobarman", []string{"bar", "foo", "the"}, []int{6, 9, 12}},
		{"abcatdogoh", []string{"cat", "dog"}, []int{2}},
		{"lingmindraboofooowingdingbarrwingmonkeypoundcake", []string{"fooo", "barr", "wing", "ding", "wing"}, []int{13}},
		{"goodgoodbestword", []string{"word", "good", "best", "good"}, []int{0}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}, []int{8}},
	}

	for _, tt := range tests {
		got := findSubstring(tt.s, tt.words)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s->%s: got %d, want %d", tt.s, tt.words, got, tt.want)
		}
	}

}
