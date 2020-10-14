package algorithms23

import (
	"reflect"
	"testing"

	"leetcode.go/lib"
)

func Test_merge(t *testing.T) {
	type args struct {
		l1 *lib.ListNode
		l2 *lib.ListNode
	}
	tests := []struct {
		name string
		args args
		want *lib.ListNode
	}{
		// TODO: Add test cases.
		{"1,3,5 - 2,4->1,2,3,4,5", args{lib.ArrayToListNode([]int{1, 3, 5}), lib.ArrayToListNode([]int{2, 4})}, lib.ArrayToListNode([]int{1, 2, 3, 4, 5})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*lib.ListNode
	}
	tests := []struct {
		name string
		args args
		want *lib.ListNode
	}{
		// TODO: Add test cases.
		{"[1,4,5],[1,3,4],[2,6]->[1,1,2,3,4,4,5,6]", args{[]*lib.ListNode{
			lib.ArrayToListNode([]int{1, 4, 5}),
			lib.ArrayToListNode([]int{1, 3, 4}),
			lib.ArrayToListNode([]int{2, 6}),
		}}, lib.ArrayToListNode([]int{1, 1, 2, 3, 4, 4, 5, 6})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeKListsV2(t *testing.T) {
	type args struct {
		lists []*lib.ListNode
	}
	tests := []struct {
		name string
		args args
		want *lib.ListNode
	}{
		// TODO: Add test cases.
		{"[1,4,5],[1,3,4],[2,6]->[1,1,2,3,4,4,5,6]", args{[]*lib.ListNode{
			lib.ArrayToListNode([]int{1, 4, 5}),
			lib.ArrayToListNode([]int{1, 3, 4}),
			lib.ArrayToListNode([]int{2, 6}),
		}}, lib.ArrayToListNode([]int{1, 1, 2, 3, 4, 4, 5, 6})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKListsV2(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKListsV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
