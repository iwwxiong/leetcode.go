package algorithms83

import (
	"reflect"
	"testing"

	"leetcode.go/lib"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"112->12", args{lib.ArrayToListNode([]int{1, 1, 2})}, lib.ArrayToListNode([]int{1, 2})},
		{"11233->123", args{lib.ArrayToListNode([]int{1, 1, 2, 3, 3})}, lib.ArrayToListNode([]int{1, 2, 3})},
		{"000->0", args{lib.ArrayToListNode([]int{0, 0, 0})}, lib.ArrayToListNode([]int{0})},
		{"-1-2-2001123->-1-20123", args{lib.ArrayToListNode([]int{-1, -2, -2, 0, 0, 1, 1, 2, 3})}, lib.ArrayToListNode([]int{-1, -2, 0, 1, 2, 3})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
