package groupanagrams

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{"example1",
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{"example2", []string{""}, [][]string{{""}}},
		{"example3", []string{"a"}, [][]string{{"a"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// tests can fail because of the order of the string in slices
			if got := groupAnagrams(tt.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
