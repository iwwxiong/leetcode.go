package algorithms62

import (
	"testing"
)

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"3,7->28", args{3, 7}, 28},
		{"3,2->3", args{3, 2}, 3},
		{"7,3->28", args{7, 3}, 28},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uniquePathsV1(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"3,7->28", args{3, 7}, 28},
		{"3,2->3", args{3, 2}, 3},
		{"7,3->28", args{7, 3}, 28},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsV1(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePathsV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
