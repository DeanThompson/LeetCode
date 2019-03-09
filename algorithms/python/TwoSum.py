#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
https://oj.leetcode.com/problems/two-sum/
"""


class Solution:
    # @return a tuple, (index1, index2)
    def twoSum(self, num, target):
        d = {}
        for idx, val in enumerate(num):
            v = d.get(val, [])
            v.append(idx + 1)
            d[val] = v

        index1, index2 = 1, 2

        for val, idx in d.items():
            idx2 = d.get(target - val)
            if idx2 is None:
                continue

            index1 = idx[0]
            index2 = idx2[0]
            if idx2 == idx:
                if len(idx2) == 1:
                    continue
                else:
                    index2 = idx2[1]
            break

        return tuple(sorted([index1, index2]))


if __name__ == '__main__':
    s = Solution()
    print(s.twoSum([0, 4, 3, 0], 0))
    print(s.twoSum([3, 2, 4], 6))
