package main

// 700. Search in a Binary Search Tree
// https://leetcode.com/problems/search-in-a-binary-search-tree/
//
// Given the root node of a binary search tree (BST) and a value.
// You need to find the node in the BST that the node's value equals the given value.
// Return the subtree rooted with that node. If such node doesn't exist, you should return NULL.

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	return searchBSTIterate(root, val)
}

func searchBSTRecursive(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}

	if root.Val > val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)
}

func searchBSTIterate(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	for root != nil {
		if root.Val == val {
			break
		}

		if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}
