package algorithms2

import (
	"reflect"
	"testing"

	"github.com/leetcode.go/lib"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"0+0=0", args{lib.ArrayToListNode([]int{0}), lib.ArrayToListNode([]int{0})}, lib.ArrayToListNode([]int{0})},
		{"123+456=579", args{lib.ArrayToListNode([]int{1, 2, 3}), lib.ArrayToListNode([]int{4, 5, 6})}, lib.ArrayToListNode([]int{5, 7, 9})},
		{"123+596=6101", args{lib.ArrayToListNode([]int{1, 2, 3}), lib.ArrayToListNode([]int{5, 9, 6})}, lib.ArrayToListNode([]int{6, 1, 0, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var want = lib.ListNodeToArray(tt.want)
			if got := lib.ListNodeToArray(addTwoNumbers(tt.args.l1, tt.args.l2)); !reflect.DeepEqual(got, want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
