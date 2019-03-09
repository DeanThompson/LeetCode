package main

import "fmt"

// 965. Univalued Binary Tree
// https://leetcode.com/problems/univalued-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	for _, nums := range [][]int{
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 2},
	} {
		tree := buildBST(nums)
		fmt.Println(nums, isUnivalTree(tree))
	}
}

func buildBST(nums []int) *TreeNode {
	var root *TreeNode
	for _, v := range nums {
		root = insertIntoBST(root, v)
	}
	return root
}

func insertIntoBST(root *TreeNode, value int) *TreeNode {
	return insertIntoBSTRecursive(root, value)
}

func insertIntoBSTRecursive(root *TreeNode, value int) *TreeNode {
	if root == nil {
		root = &TreeNode{value, nil, nil}
	} else if value < root.Val {
		root.Left = insertIntoBSTRecursive(root.Left, value)
	} else {
		root.Right = insertIntoBSTRecursive(root.Right, value)
	}
	return root
}

func isUnivalTree(root *TreeNode) bool {
	return traverse(root, root.Val)
}

func traverse(root *TreeNode, n int) bool {
	if root == nil {
		return true
	}
	if root.Val != n {
		return false
	}

	if !traverse(root.Left, n) {
		return false
	}
	if !traverse(root.Right, n) {
		return false
	}
	return true
}
