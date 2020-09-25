package algorithms100

import "leetcode.go/lib"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode lib.TreeNode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}
	var v bool = p.Val == q.Val
	var left bool = isSameTree(p.Left, q.Left)
	var right bool = isSameTree(p.Right, q.Right)
	return v && left && right
}
