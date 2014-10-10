#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
https://oj.leetcode.com/problems/binary-tree-preorder-traversal/
"""

# Definition for a  binary tree node
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    # @param root, a tree node
    # @return a list of integers
    def preorderTraversal(self, root):
        result = []
        self._traversal(root, lambda x: result.append(x))
        return result

    def _traversal(self, root, callback):
        if root is None:
            return

        callback(root.val)
        self._traversal(root.left, callback)
        self._traversal(root.right, callback)


if __name__ == '__main__':
    s = Solution()
    root = TreeNode(1)
    root.right = TreeNode(2)
    root.right.left = TreeNode(3)
    print s.preorderTraversal(root)

