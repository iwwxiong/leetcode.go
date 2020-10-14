package algorithms24

import (
	"leetcode.go/lib"
)

func swapPairs(head *lib.ListNode) *lib.ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var result *lib.ListNode = new(lib.ListNode)
	var cur *lib.ListNode = result
	for head != nil {
		if head.Next == nil {
			cur.Val = head.Val
			head = head.Next
			if head != nil {
				cur.Next = new(lib.ListNode)
				cur = cur.Next
			}
		} else {
			cur.Next = new(lib.ListNode)
			cur.Val = head.Next.Val
			cur = cur.Next

			cur.Val = head.Val

			head = head.Next.Next
			if head != nil {
				cur.Next = new(lib.ListNode)
				cur = cur.Next
			}

		}
	}
	return result
}
