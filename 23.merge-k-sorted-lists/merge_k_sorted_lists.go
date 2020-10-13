package algorithms23

import (
	"leetcode.go/lib"
)

func merge(l1 *lib.ListNode, l2 *lib.ListNode) *lib.ListNode {
	var l *lib.ListNode = new(lib.ListNode)
	var cur *lib.ListNode = l
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = new(lib.ListNode)
			cur.Next.Val = l1.Val
			cur = cur.Next
			l1 = l1.Next
		} else {
			cur.Next = new(lib.ListNode)
			cur.Next.Val = l2.Val
			cur = cur.Next
			l2 = l2.Next
		}
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return l.Next
}

func mergeKLists(lists []*lib.ListNode) *lib.ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	var head *lib.ListNode = merge(lists[0], lists[1])
	for i := 2; i <= len(lists)-1; i++ {
		head = merge(head, lists[i])
	}
	return head
}
