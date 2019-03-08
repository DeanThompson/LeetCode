package main

import "fmt"

// 938. Range Sum of BST
// https://leetcode.com/problems/range-sum-of-bst/
//
// Given the root node of a binary search tree, return the sum of values of all nodes with value between L and R (inclusive).
//
// The binary search tree is guaranteed to have unique values.
//
// 相关资料
// 维基百科：二叉搜索树 https://zh.wikipedia.org/wiki/%E4%BA%8C%E5%85%83%E6%90%9C%E5%B0%8B%E6%A8%B9
// Binary Search Tree (BST) 具有以下特性
//   - 若任意节点的左子树不空，则左子树上所有节点的值均小于它的根节点的值；
//   - 若任意节点的右子树不空，则右子树上所有节点的值均大于它的根节点的值；
//   - 任意节点的左、右子树也分别为二叉查找树；
//   - 没有键值相等的节点。
// 二叉查找树相比于其他数据结构的优势在于查找、插入的时间复杂度较低。为 O(log n)
//
// 题目是要求一个区间里所有节点的和。根据 BST 的特点，如果有比根结点小的，则要继续查找左子树；更大的就是右子树。
// 因此这就是一个深搜问题，用递归实现。

func main() {
	nums := []int{10, 5, 15, 3, 7, -1, 18}
	L, R := 7, 15
	tree := buildBST(nums)
	printBST(tree)
	rv := rangeSumBST(tree, L, R)
	fmt.Println(rv)
}

func buildBST(nums []int) *TreeNode {
	var root *TreeNode
	for _, v := range nums {
		root = insertBST(root, v)
	}
	return root
}

func insertBST(root *TreeNode, value int) *TreeNode {
	if root == nil {
		root = &TreeNode{value, nil, nil}
	} else if value < root.Val {
		root.Left = insertBST(root.Left, value)
	} else {
		root.Right = insertBST(root.Right, value)
	}
	return root
}

func printBST(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	printBST(root.Left)
	printBST(root.Right)
}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	result := 0
	dfs(root, L, R, &result)
	return result
}

func dfs(root *TreeNode, L int, R int, result *int) {
	if root == nil {
		return
	}
	if L <= root.Val && root.Val <= R {
		*result += root.Val
	}
	if L < root.Val {
		dfs(root.Left, L, R, result)
	}
	if root.Val < R {
		dfs(root.Right, L, R, result)
	}
}
