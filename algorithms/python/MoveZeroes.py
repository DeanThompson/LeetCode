# -*- coding: utf-8 -*-

__author__ = 'leon'

"""
https://leetcode.com/problems/move-zeroes/
"""


class Solution(object):
    def moveZeroes(self, nums):
        """
        :type nums: List[int]
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        index, count = 0, 0
        while count < len(nums):
            if nums[index] == 0:
                nums.pop(index)
                nums.append(0)
                index -= 1
            index += 1
            count += 1
