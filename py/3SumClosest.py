#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
https://oj.leetcode.com/problems/3sum-closest/
"""


class Solution:
    # @return an integer
    def threeSumClosest(self, num, target):
        min_diff = 1<<31
        result = 0

        num.sort()

        length = len(num)
        for i, v in enumerate(num):
            j = i + 1
            k = length - 1
            while j < k:
                total = v + num[j] + num[k]
                diff = abs(total-target)
                if diff < min_diff:
                    if diff == 0:
                        return total
                    min_diff = diff
                    result = total
                if total < target:
                    j += 1
                else:
                    k -= 1
        
        return result


if __name__ == '__main__':
    s = Solution()
    print s.threeSumClosest([-1, 2, 1, -4], 1)

