package algorithms19

import (
	"reflect"
	"testing"

	"leetcode.go/lib"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *lib.ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *lib.ListNode
	}{
		// TODO: Add test cases.
		{"[1],1->[]", args{lib.ArrayToListNode([]int{1}), 1}, lib.ArrayToListNode([]int{})},
		{"[1,2],2->[2]", args{lib.ArrayToListNode([]int{1, 2}), 2}, lib.ArrayToListNode([]int{2})},
		{"[1,2,3,4,5],2->[1,2,3,5]", args{lib.ArrayToListNode([]int{1, 2, 3, 4, 5}), 2}, lib.ArrayToListNode([]int{1, 2, 3, 5})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
