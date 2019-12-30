package algorithms11

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"[2,3,10,5,7,8,9]->36", args{[]int{2, 3, 10, 5, 7, 8, 9}}, 36},
		{"[1,8,6,2,5,4,8,3,7]->49", args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, 49},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
