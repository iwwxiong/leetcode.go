package algorithms66

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"[0]->[1]", args{[]int{0}}, []int{1}},
		{"[4,3,2,1]->[4,3,2,2]", args{[]int{4, 3, 2, 1}}, []int{4, 3, 2, 2}},
		{"[9,9,9]->[1,0,0,0]", args{[]int{9, 9, 9}}, []int{1, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
