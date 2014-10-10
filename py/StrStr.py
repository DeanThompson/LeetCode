#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
https://oj.leetcode.com/problems/implement-strstr/
"""


class Solution:
    # @param haystack, a string
    # @param needle, a string
    # @return a string or None
    def strStr(self, haystack, needle):
        length = len(needle)
        for idx in xrange(len(haystack)-length+1):
            # print haystack[idx:(idx+length)]
            if haystack[idx:(idx+length)] == needle:
                return haystack[idx:]
        return None

    def strStr2(self, haystack, needle):
        try:
            return haystack[haystack.index(needle):]
        except:
            return None

    # TODO: implenment using KMP algorithm
    def strStr3(self, haystack, needle):
        pass


if __name__ == '__main__':
    s = Solution()
    haystack = "this is a test"
    needle = "test"
    print s.strStr(haystack, needle)
    print s.strStr2(haystack, needle)

    needle = "is a"
    print s.strStr(haystack, needle)
    print s.strStr2(haystack, needle)
