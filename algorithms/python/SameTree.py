#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
https://oj.leetcode.com/problems/same-tree/
"""


# Definition for a  binary tree node
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    # @param p, a tree node
    # @param q, a tree node
    # @return a boolean
    def isSameTree(self, p, q):
        if p is None or q is None:
            return p == q
        return p.val == q.val and self.isSameTree(p.left, q.left) and self.isSameTree(p.right, q.right)


if __name__ == '__main__':
    proot = TreeNode(1)
    proot.left = TreeNode(2)
    proot.right = TreeNode(3)

    qroot = TreeNode(1)
    qroot.left = TreeNode(2)
    qroot.right = TreeNode(3)

    s = Solution()
    print(s.isSameTree(proot, qroot))
