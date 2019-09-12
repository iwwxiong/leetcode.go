package algorithms7

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"0->0", args{0}, 0},
		{"123->321", args{123}, 321},
		{"-123->-321", args{-123}, -321},
		{"120->21", args{120}, 21},
		{"-654->-456", args{-654}, -456},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
