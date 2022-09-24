package minimumwindowsubstring

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"example_1", args{"ADOBECODEBANC", "ABC"}, "BANC"},
		{"example_2", args{"a", "a"}, "a"},
		{"example_3", args{"a", "aa"}, ""},
		{"example_3", args{"ab", "b"}, "b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
