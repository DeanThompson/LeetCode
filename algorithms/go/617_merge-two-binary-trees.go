package main

// 617. Merge Two Binary Trees
// https://leetcode.com/problems/merge-two-binary-trees/
// Given two binary trees and imagine that when you put one of them to cover the other,
// some nodes of the two trees are overlapped while the others are not.
//
//You need to merge them into a new binary tree. The merge rule is that if two nodes overlap,
// then sum node values up as the new value of the merged node. Otherwise,
// the NOT null node will be used as the node of new tree.

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	node := &TreeNode{}

	if t1 == nil {
		node.Val = t2.Val
		node.Left = mergeTrees(nil, t2.Left)
		node.Right = mergeTrees(nil, t2.Right)
	} else if t2 == nil {
		node.Val = t1.Val
		node.Left = mergeTrees(t1.Left, nil)
		node.Right = mergeTrees(t1.Right, nil)
	} else {
		node.Val = t1.Val + t2.Val
		node.Left = mergeTrees(t1.Left, t2.Left)
		node.Right = mergeTrees(t1.Right, t2.Right)
	}
	return node
}
