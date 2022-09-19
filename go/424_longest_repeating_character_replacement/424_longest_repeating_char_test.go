package longestrepeatingcharacterreplacement

import "testing"

func Test_characterReplacement(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{"ABAB", 2}, 4},    // => [A(A)A(A)]
		{"example2", args{"AABABBA", 1}, 4}, // => [AA(A)A]BBA
		{"example3", args{"AABA", 0}, 2},    // => [AA(A)A]BBA
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := characterReplacement(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("characterReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
