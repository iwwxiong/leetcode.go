package algorithms15

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"{-1, 0, 1, 2, -1, -4}->{{-1, 0, 1}, {-1, -1, 2}}", args{[]int{-1, 0, 1, 2, -1, -4}}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sorted(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"{5,3,2,4,5,1}->{1,2,3,4,5,5}", args{[]int{5, 3, 2, 4, 5, 1}}, []int{1, 2, 3, 4, 5, 5}},
		{"{-1,0,-2,3,2}->{-2,-1,0,2,3}", args{[]int{-1, 0, -2, 3, 2}}, []int{-2, -1, 0, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sorted(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
