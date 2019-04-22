package main

// https://leetcode.com/problems/path-sum/
//
// 112. Path Sum
//
// Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum.
//
// Note: A leaf is a node with no children.

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && sum == root.Val {
		return true
	}
	remain := sum - root.Val
	return hasPathSum(root.Left, remain) || hasPathSum(root.Right, remain)
}
