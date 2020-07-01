package algorithms21

import (
	"leetcode.go/lib"
)

// ListNode Definition for singly-linked list.
type ListNode = lib.ListNode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var result = &ListNode{}
	curr := result

	for l1 != nil {
		for l2 != nil && l1.Val > l2.Val {
			curr.Next = new(ListNode)
			curr.Next.Val = l2.Val
			curr = curr.Next
			l2 = l2.Next
		}
		curr.Next = new(ListNode)
		curr.Next.Val = l1.Val
		curr = curr.Next
		l1 = l1.Next
	}
	if l2 != nil {
		curr.Next = l2
	}
	return result.Next
}
