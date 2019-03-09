#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
https://oj.leetcode.com/problems/path-sum/
"""


# Definition for a  binary tree node
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

    def __repr__(self):
        return str(self.val)


class Solution:
    # @param root, a tree node
    # @param sum, an integer
    # @return a boolean
    def hasPathSum(self, root, sum):
        if root is None:
            return False
        stack = [root]
        total = 0
        while len(stack) > 0:
            node = stack.pop()
            total += node.val

            leaf = True
            if node.right is not None:
                stack.append(node.right)
                leaf = False
            if node.left is not None:
                stack.append(node.left)
                leaf = False
            if leaf:
                if total == sum:
                    return True
                total -= node.val

        return False


if __name__ == '__main__':
    s = Solution()
    root = TreeNode(5)

    root.left = TreeNode(4)
    root.right = TreeNode(8)

    root.left.left = TreeNode(11)
    root.right.left = TreeNode(13)
    root.right.right = TreeNode(4)

    root.left.left.left = TreeNode(7)
    root.left.left.right = TreeNode(2)
    root.right.right.right = TreeNode(1)

    print(s.hasPathSum(root, 22))

    root = TreeNode(1)
    root.left = TreeNode(2)

    print(s.hasPathSum(root, 1))
