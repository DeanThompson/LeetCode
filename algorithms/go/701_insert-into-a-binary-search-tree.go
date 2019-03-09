package main

import "fmt"

// 701. Insert into a Binary Search Tree
// https://leetcode.com/problems/insert-into-a-binary-search-tree/

func main() {
	nums := []int{1, 2, 3, 4, 7}
	tree := buildBST(nums)
	printBST(tree)
	tree = insertIntoBST(tree, 5)
	printBST(tree)
}

func buildBST(nums []int) *TreeNode {
	var root *TreeNode
	for _, v := range nums {
		root = insertIntoBST(root, v)
	}
	return root
}

func insertIntoBST(root *TreeNode, value int) *TreeNode {
	return insertIntoBSTIter(root, value)
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

func insertIntoBSTIter(root *TreeNode, value int) *TreeNode {
	node := root
	for {
		if value < node.Val {
			if node.Left != nil {
				node = node.Left
			} else {
				node.Left = &TreeNode{value, nil, nil}
				break
			}
		} else {
			if node.Right != nil {
				node = node.Right
			} else {
				node.Right = &TreeNode{value, nil, nil}
				break
			}
		}
	}
	return root
}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printBST(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	printBST(root.Left)
	printBST(root.Right)
}
