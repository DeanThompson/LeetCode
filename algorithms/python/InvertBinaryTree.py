# -*- coding: utf-8 -*-

__author__ = 'leon'

"""
https://leetcode.com/problems/invert-binary-tree/
"""


class Solution(object):
    def invertTree(self, root):
        """
        :type root: TreeNode
        :rtype: TreeNode
        """
        if not root:
            return root
        temp = root.left
        root.left = self.invertTree(root.right)
        root.right = self.invertTree(temp)
        return root
