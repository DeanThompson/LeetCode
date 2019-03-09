#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
https://oj.leetcode.com/problems/min-stack/
"""


class MinStack:
    def __init__(self):
        self._stack = []
        self._min_index = []

    # @param x, an integer
    # @return an integer
    def push(self, x):
        try:
            if x < self.getMin():
                self._min_index.append(len(self._stack))
        except:
            self._min_index.append(0)
        self._stack.append(x)

    # @return nothing
    def pop(self):
        self._stack.pop()
        if self._min_index[-1] >= len(self._stack):
            self._min_index.pop()

    # @return an integer
    def top(self):
        return self._stack[-1]

    # @return an integer
    def getMin(self):
        return self._stack[self._min_index[-1]]

    def __str__(self):
        return "%s, %s" % (self._stack, self._min_index)


if __name__ == '__main__':
    st = MinStack()
    st.push(100)
    print(st)
    st.push(1)
    print(st)
    st.push(34)
    print(st)

    print(st.getMin())
