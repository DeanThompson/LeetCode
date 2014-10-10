#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
https://oj.leetcode.com/problems/maximum-depth-of-binary-tree/
"""


# Definition for a  binary tree node
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    # @param root, a tree node
    # @return an integer
    def maxDepth(self, root):
        if root is None:
            return 0

        left_depth = self.maxDepth(root.left)
        right_depth = self.maxDepth(root.right)
        return max(left_depth, right_depth) + 1


if __name__ == '__main__':
    s = Solution()
    root = TreeNode(1)
    print s.maxDepth(root)

    root.left = TreeNode(2)
    root.right = TreeNode(3)
    print s.maxDepth(root)


    root.left.left = TreeNode(4)
    root.left.left.right = TreeNode(4)
    print s.maxDepth(root)
