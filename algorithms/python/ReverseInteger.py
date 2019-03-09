#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
https://oj.leetcode.com/problems/reverse-integer/
"""


class Solution:

    # @return an integer
    def reverse(self, x):
        factor = 1 if x > 0 else -1
        return factor * int(str(abs(x))[::-1])

    # TODO: implement without using `str`
    def reverse2(self, x):
        pass


if __name__ == '__main__':
    import time


    def timeit(func):
        def wrapper(*args, **kwargs):
            loop = 1000000
            i = 0
            start = time.time()
            while i < loop:
                func(*args, **kwargs)
                i += 1
            duration = time.time() - start
            avg = duration / loop
            print("duration:", duration, "loop:", loop, "avg:", avg)

        return wrapper


    s = Solution()


    @timeit
    def run():
        s.reverse(123462841425)


    run()
