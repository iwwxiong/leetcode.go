package lib

import (
	"reflect"
	"testing"
)

func TestArrayToListNode(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayToListNode(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayToListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListNodeToArray(t *testing.T) {
	type args struct {
		ln *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListNodeToArray(tt.args.ln); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListNodeToArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseListNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"[1,2,3]->[3,2,1]", args{ArrayToListNode([]int{1, 2, 3})}, ArrayToListNode([]int{3, 2, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseListNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTailNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"[1,2,3]->[3]", args{ArrayToListNode([]int{1, 2, 3})}, ArrayToListNode([]int{3})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TailNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TailNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
