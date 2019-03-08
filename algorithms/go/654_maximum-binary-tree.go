package main

import "fmt"

// 654. Maximum Binary Tree
// https://leetcode.com/problems/maximum-binary-tree/

func main() {
	nums := []int{3, 2, 1, 6, 0, 5}
	t := constructMaximumBinaryTree(nums)
	fmt.Println(t.Val)
}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路：首先找出最大值和对应的下标，然后递归构建左右子树
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxIdx := 0
	max := nums[maxIdx]
	for i := maxIdx + 1; i < len(nums); i++ {
		if nums[i] > max {
			maxIdx = i
			max = nums[i]
		}
	}

	root := &TreeNode{
		Val:   max,
		Left:  constructMaximumBinaryTree(nums[0:maxIdx]),
		Right: constructMaximumBinaryTree(nums[maxIdx+1:]),
	}
	return root
}
