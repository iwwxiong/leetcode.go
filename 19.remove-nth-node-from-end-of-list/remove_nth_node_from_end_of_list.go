package algorithms19

import (
	"leetcode.go/lib"
)

func removeNthFromEnd(head *lib.ListNode, n int) *lib.ListNode {
	var h *lib.ListNode = head
	var length int
	for h != nil {
		h = h.Next
		length++
	}
	if length == 0 {
		return head.Next
	}
	if length == 1 {
		return &lib.ListNode{}
	}
	if length == n {
		return head.Next
	}
	var result *lib.ListNode = head
	for i := 0; i < length-n-1; i++ {
		result = result.Next
	}
	result.Next = result.Next.Next
	return head
}
