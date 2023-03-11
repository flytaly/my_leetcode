package lengthOfLastWord

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
	}
	for _, tt := range tests {
		if got := lengthOfLastWord(tt.s); got != tt.want {
			t.Errorf("lengthOfLastWord(%q) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
