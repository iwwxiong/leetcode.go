package algorithms18

import (
	"reflect"
	"testing"
)

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"[],0->[]", args{[]int{}, 0}, [][]int{}},
		{"[1,0,-1,0,-2,2],0->[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]", args{[]int{1, 0, -1, 0, -2, 2}, 0}, [][]int{[]int{-2, -1, 1, 2}, []int{-2, 0, 0, 2}, []int{-1, 0, 0, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
