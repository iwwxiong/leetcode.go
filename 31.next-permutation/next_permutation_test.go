package algorithms31

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"[3,2,1]->[1,2,3]", args{[]int{3, 2, 1}}, []int{1, 2, 3}},
		{"[1,2,3]->[1,3,2]", args{[]int{1, 2, 3}}, []int{1, 3, 2}},
		{"[1,5,8,4,7,6,5,3,1]->[1,5,8,5,1,3,4,6,7]", args{[]int{1, 5, 8, 4, 7, 6, 5, 3, 1}}, []int{1, 5, 8, 5, 1, 3, 4, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextPermutation(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
