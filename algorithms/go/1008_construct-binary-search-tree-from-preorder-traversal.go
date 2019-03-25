package main

// 1008. Construct Binary Search Tree from Preorder Traversal
// https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/

// Return the root node of a binary search tree that matches the given preorder traversal.
//
// (Recall that a binary search tree is a binary tree where for every node,
// any descendant of node.left has a value < node.val, and any descendant of node.
// right has a value > node.val.
// Also recall that a preorder traversal displays the value of the node first,
// then traverses node.left, then traverses node.right.)

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	root := &TreeNode{preorder[0], nil, nil}
	for i := 1; i < len(preorder); i++ {
		insertBstDfs(root, preorder[i])
	}
	return root
}

func insertBstDfs(root *TreeNode, value int) {
	node := &TreeNode{value, nil, nil}
	if value < root.Val {
		if root.Left == nil {
			root.Left = node
		} else {
			insertBstDfs(root.Left, value)
		}
	} else {
		if root.Right == nil {
			root.Right = node
		} else {
			insertBstDfs(root.Right, value)
		}
	}
}
