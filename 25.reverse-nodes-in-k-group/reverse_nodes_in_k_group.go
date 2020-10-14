package algorithms25

import (
	"fmt"

	"leetcode.go/lib"
)

// 翻转链表
func reverse(head *lib.ListNode) *lib.ListNode {
	var s *lib.ListNode
	for head != nil {
		var next *lib.ListNode = head.Next
		head.Next = s
		s = head
		head = next
	}
	return s
}

func reverseKGroup(head *lib.ListNode, k int) *lib.ListNode {
	if k == 1 {
		return head
	}
	var i int = 1
	var result *lib.ListNode = new(lib.ListNode)
	var temp *lib.ListNode = new(lib.ListNode)
	var cur *lib.ListNode = temp
	for head != nil {
		cur.Val = head.Val
		head = head.Next
		if i < k {
			cur.Next = new(lib.ListNode)
			cur = cur.Next
		}

		if i == k {
			i = 1
			// fmt.Println(lib.ListNodeToArray(temp))
			temp = reverse(temp)
			result.Next = temp
		} else {
			i++
		}

	}
	// fmt.Println(lib.ListNodeToArray(temp))
	if temp != nil {
		result.Next = temp
	}
	fmt.Println(lib.ListNodeToArray(result.Next))
	return result.Next
}
