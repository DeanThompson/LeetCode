#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
https://oj.leetcode.com/problems/binary-tree-inorder-traversal/
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
    def inorderTraversal(self, root):
        result = []
        self._traversal(root, lambda x: result.append(x))
        return result

    def _traversal(self, root, callback):
        if root is None:
            return
        self._traversal(root.left, callback)
        callback(root.val)
        self._traversal(root.right, callback)


if __name__ == '__main__':
    s = Solution()
    root = TreeNode(1)
    root.right = TreeNode(2)
    root.right.left = TreeNode(3)
    print(s.inorderTraversal(root))

    root = TreeNode(1)
    root.left = TreeNode(2)
    root.right = TreeNode(3)
    root.right.left = TreeNode(4)
    root.right.left.right = TreeNode(5)
    print(s.inorderTraversal(root))
