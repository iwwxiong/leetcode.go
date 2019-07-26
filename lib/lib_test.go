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
