#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
https://oj.leetcode.com/problems/reverse-words-in-a-string/
"""


class Solution:
    # @param s, a string
    # @return a string
    def reverseWords(self, s):
        return " ".join(reversed(filter(None, s.split(" "))))

    def reverseWords2(self, s):
        if s == "":
            return ""

        spaces = (" ", "\t", "\r", "\n")
        words = []
        append = words.append
        i = len(s)
        while i > 0:
            while i > 0 and s[i - 1] in spaces:
                i -= 1
            if i == 0:
                break
            right = i
            while i > 0 and s[i - 1] not in spaces:
                i -= 1
            left = i
            append(s[left:right])

        return " ".join(words)


if __name__ == '__main__':
    ss = Solution()
    s = "the \tsky is         blue"
    print(ss.reverseWords(s))

    print(ss.reverseWords2(s))
