# -*- coding: utf-8 -*-

"""
https://leetcode.com/problems/add-digits/
"""


class Solution(object):
    def addDigits(self, num):
        """
        :type num: int
        :rtype: int
        """
        if num < 10:
            return num
        return 1 + (num - 1) % 9

    def addDigits_recursive(self, num):
        """
        :type num: int
        :rtype: int
        """
        if num < 10:
            return num
        return self.addDigits(sum(map(int, str(num))))

    def addDigits_loop(self, num):
        """
        :type num: int
        :rtype: int
        """
        if num < 10:
            return num
        while True:
            pre, digit = divmod(num, 10)
            total = digit
            while pre > 10:
                pre, digit = divmod(pre, 10)
                total += digit
            total += pre
            if total < 10:
                return total
            else:
                num = total
