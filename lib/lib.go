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

// SortIntSlice sort int slice
func SortIntSlice(nums []int) []int {
	var left, right []int
	if len(nums) <= 1 {
		return nums
	}
	var k int = nums[0]
	for _, num := range nums[1:] {
		if num < k {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}
	left = SortIntSlice(left)
	right = SortIntSlice(right)
	left = append(left, k)
	left = append(left, right...)
	return left
}

// 翻转链表
func ReverseListNode(head *ListNode) *ListNode {
	var tail *ListNode
	for head != nil {
		var next *ListNode = head.Next
		head.Next = tail
		tail = head
		head = next
	}
	return tail
}

// 获取单链表尾节点
func TailNode(head *ListNode) *ListNode {
	for head.Next != nil {
		head = head.Next
	}
	return head
}
