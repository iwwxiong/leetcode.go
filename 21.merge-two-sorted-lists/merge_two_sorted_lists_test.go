package algorithms21

import (
	"reflect"
	"testing"

	"leetcode.go/lib"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"0,0=0,0", args{lib.ArrayToListNode([]int{0}), lib.ArrayToListNode([]int{0})}, lib.ArrayToListNode([]int{0, 0})},
		{"123,456=123456", args{lib.ArrayToListNode([]int{1, 2, 3}), lib.ArrayToListNode([]int{4, 5, 6})}, lib.ArrayToListNode([]int{1, 2, 3, 4, 5, 6})},
		{"456+12345=6101", args{lib.ArrayToListNode([]int{4, 5, 6}), lib.ArrayToListNode([]int{1, 2, 3, 4, 5})}, lib.ArrayToListNode([]int{1, 2, 3, 4, 4, 5, 5, 6})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
