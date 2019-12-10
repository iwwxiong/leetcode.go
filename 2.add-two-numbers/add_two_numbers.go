package algorithms2

import (
	"leetcode.go/lib"
)

// ListNode Definition for singly-linked list.
type ListNode = lib.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result = &ListNode{}
	cur := result
	var decimal = 0

	for l1 != nil || l2 != nil {
		var sum = 0
		sum += decimal

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		decimal = sum / 10
		cur.Next = new(ListNode)
		cur.Next.Val = sum % 10
		cur = cur.Next

		if l1 == nil && l2 == nil && decimal > 0 {
			cur.Next = new(ListNode)
			cur.Next.Val = decimal
		}
	}
	return result.Next
}

