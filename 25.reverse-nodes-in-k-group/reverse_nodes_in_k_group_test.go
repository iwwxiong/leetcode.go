package algorithms25

import (
	"reflect"
	"testing"

	"leetcode.go/lib"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *lib.ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *lib.ListNode
	}{
		// TODO: Add test cases.
		{"[1,2,3,4,5],3->[3,2,1,4,5]", args{lib.ArrayToListNode([]int{1, 2, 3, 4, 5}), 3}, lib.ArrayToListNode([]int{3, 2, 1, 4, 5})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
