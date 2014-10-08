#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
https://oj.leetcode.com/problems/single-number/
"""


class Solution:
    # @param A, a list of integer
    # @return an integer
    def singleNumber(self, A):
        return reduce(lambda x, y: x^y, A)


if __name__ == '__main__':
    s = Solution()
    print s.singleNumber([1, 1, 2, 2, 3])
    print s.singleNumber([4, 5, 6, 7, 8, 8, 7, 6, 5])


