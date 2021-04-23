package algorithms72

import (
	"testing"
)

func Test_minDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"horse,ros->3", args{"horse", "ros"}, 3},
		{"intention,execution->5", args{"intention", "execution"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min3int(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1,2,3->1", args{1, 2, 3}, 1},
		{"1,3,2->1", args{1, 3, 2}, 1},
		{"3,2,1->1", args{3, 2, 1}, 1},
		{"1,1,2->1", args{1, 1, 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min3int(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("min3int() = %v, want %v", got, tt.want)
			}
		})
	}
}
