package courseschedule

import "testing"

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{2, [][]int{{1, 0}}}, true},
		{"Example 2", args{2, [][]int{{1, 0}, {0, 1}}}, false},
		{"Example 3", args{5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
