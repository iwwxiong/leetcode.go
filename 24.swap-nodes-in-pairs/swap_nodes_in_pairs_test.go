package algorithms24

import (
	"reflect"
	"testing"

	"leetcode.go/lib"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *lib.ListNode
	}
	tests := []struct {
		name string
		args args
		want *lib.ListNode
	}{
		// TODO: Add test cases.
		{"[]->[]", args{nil}, nil},
		{"[1]->[1]", args{lib.ArrayToListNode([]int{1})}, lib.ArrayToListNode([]int{1})},
		{"[1,2,3,4]->[2,1,4,3]", args{lib.ArrayToListNode([]int{1, 2, 3, 4})}, lib.ArrayToListNode([]int{2, 1, 4, 3})},
		{"[1,2,3,4,5]->[2,1,4,3,5]", args{lib.ArrayToListNode([]int{1, 2, 3, 4, 5})}, lib.ArrayToListNode([]int{2, 1, 4, 3, 5})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
