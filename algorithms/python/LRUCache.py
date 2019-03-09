#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
https://oj.leetcode.com/problems/lru-cache/

inspired by the implementation from jinja2
"""


# TODO: `list.remove` is too slow
class LRUCache:

    # @param capacity, an integer
    def __init__(self, capacity):
        self._cap = capacity
        self._cache = {}
        self._queue = []

    # @return an integer
    def get(self, key):
        value = self._cache.get(key)
        if value is None:
            return -1

        # move to head
        if self._queue[-1] != key:
            self._queue.remove(key)
            self._queue.append(key)

        return value

    # @param key, an integer
    # @param value, an integer
    # @return nothing
    def set(self, key, value):
        if key in self._cache:
            self._queue.remove(key)
        else:
            if len(self._cache) == self._cap:
                # remove the Least Recently Used item
                rk = self._queue[0]
                del self._queue[0]
                del self._cache[rk]

        self._queue.append(key)
        self._cache[key] = value


if __name__ == '__main__':
    l = LRUCache(5)
    l.set(1, 1)
    l.set(2, 2)
    l.set(3, 3)
    l.set(4, 4)
    l.set(5, 5)
    l.set(5, 5)
    l.get(4)
    print(l._queue)
    print(l._cache)

    l.set(6, 6)
    print(l.get(2))
    print(l.get(3))
    print(l.get(100))
    print(l._queue)
    print(l._cache)
