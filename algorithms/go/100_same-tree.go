package main

// 100. Same Tree
// https://leetcode.com/problems/same-tree/

// Given two binary trees, write a function to check if they are the same or not.
//
// Two binary trees are considered the same if they are structurally identical and the nodes have the same value.

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return recursive(p, q)
}

func recursive(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && recursive(p.Left, q.Left) && recursive(p.Right, q.Right)
}
