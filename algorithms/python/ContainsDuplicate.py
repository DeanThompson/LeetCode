# -*- coding: utf-8 -*-

__author__ = 'leon'

"""
https://leetcode.com/problems/contains-duplicate/
"""


class Solution(object):
    def containsDuplicate(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        seen = set()
        for x in nums:
            if x in seen:
                return True
            seen.add(x)
        return False

    def containsDuplicate_2(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        return len(set(nums)) != len(nums)
