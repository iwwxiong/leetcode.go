package algorithms1

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"{2,7,3,8}->[0,1]", args{[]int{2, 7, 3, 8}, 9}, []int{0, 1}},
		{"{0,4,3,0}->[0,3]", args{[]int{0, 4, 3, 0}, 0}, []int{0, 3}},
		{"{2,5,0,8,1,4}->[2,3]", args{[]int{2, 5, 0, 8, 1, 4}, 8}, []int{2, 3}},
		{"{-3,4,3,90}->[0,2]", args{[]int{-3, 4, 3, 90}, 0}, []int{0, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
