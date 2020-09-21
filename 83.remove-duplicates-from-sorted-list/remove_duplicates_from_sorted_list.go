package algorithms83

import "leetcode.go/lib"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = lib.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	var result *ListNode = &ListNode{}
	var cur *ListNode = result
	var zero bool = true
	for head != nil {
		if head.Val == 0 && zero {
			cur.Next = new(ListNode)
			cur.Next.Val = head.Val
			cur = cur.Next
			zero = false
		} else if head.Val != cur.Val {
			cur.Next = new(ListNode)
			cur.Next.Val = head.Val
			cur = cur.Next
		}
		head = head.Next
	}
	return result.Next
}
