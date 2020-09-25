package lib

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ArrayToListNode int array to ListNode
func ArrayToListNode(arr []int) *ListNode {
	var head = &ListNode{}
	var cur = &ListNode{}
	head.Next = cur
	for index, num := range arr {
		cur.Val = num
		if index < len(arr)-1 {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
	}
	return head.Next
}

// ListNodeToArray ListNode to array
func ListNodeToArray(ln *ListNode) []int {
	var arr = []int{}
	for cur := ln; cur != nil; cur = cur.Next {
		arr = append(arr, cur.Val)
	}
	return arr
}
