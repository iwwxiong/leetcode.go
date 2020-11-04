package algorithms29

import (
	"testing"
)

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"10,1->10", args{10, 1}, 10},
		{"10,3->3", args{10, 3}, 3},
		{"7,-3->-2", args{7, -3}, -2},
		{"-2147483648,1->-2147483648", args{-2147483648, 1}, -2147483648},
		{"-2147483648,-1->2147483647", args{-2147483648, -1}, 2147483647},
		{"-2147483648,2->-1073741824", args{-2147483648, 2}, -1073741824},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_negativeDivide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"-7,-3->2", args{-7, -3}, 2},
		{"-9,-3->3", args{-9, -3}, 3},
		{"-10,-1->10", args{-10, -1}, 10},
		{"-10,-3->3", args{-10, -3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := negativeDivide(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("negativeDivide() = %v, want %v", got, tt.want)
			}
		})
	}
}
