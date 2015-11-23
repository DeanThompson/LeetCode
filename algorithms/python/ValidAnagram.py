# -*- coding: utf-8 -*-

__author__ = 'leon'

"""
https://leetcode.com/problems/valid-anagram/
"""


class Solution(object):
    def isAnagram_frequency_table(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        if len(s) != len(t):
            return False
        frequency_s = [0] * 26
        frequency_t = [0] * 26
        c = ord('a')
        for x in s:
            frequency_s[ord(x) - c] += 1
        for x in t:
            frequency_t[ord(x) - c] += 1
        return frequency_s == frequency_t

    def isAnagram_sort(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        if len(s) != len(t):
            return False
        return sorted(s) == sorted(t)
