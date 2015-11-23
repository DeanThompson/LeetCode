#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
https://leetcode.com/problems/nim-game/
"""


class Solution(object):
    def canWinNim(self, n):
        """
        :type n: int
        :rtype: bool
        """
        return bool(n % 4)
