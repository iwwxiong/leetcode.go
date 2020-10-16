package algorithms25

import (
	"leetcode.go/lib"
)

func reverseKGroup(head *lib.ListNode, k int) *lib.ListNode {
	if k == 1 || head == nil {
		return head
	}
	var pointer *lib.ListNode = head
	var i int = 1
	for i < k {
		pointer = pointer.Next
		if pointer == nil {
			return head
		}
		i++
	}
	var temp *lib.ListNode = pointer.Next
	// 断开链接
	pointer.Next = nil
	var result *lib.ListNode = lib.ReverseListNode(head)
	// 反转后 head 就成了 tail 节点
	head.Next = reverseKGroup(temp, k)

	return result
}
