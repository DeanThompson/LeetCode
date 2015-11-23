#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
https://oj.leetcode.com/problems/reorder-list/
"""


# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

    def __repr__(self):
        return "%s" % self.val


class Solution:
    # @param head, a ListNode
    # @return nothing
    def reorderList(self, head):
        if head is None or head.next is None:
            return
        slow = head
        fast = head
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next

        right_head = slow.next     # slow is currently middle
        slow.next = None
        right_head = self._reverse(right_head)  # reverse right part

        # merge left part and the reversed right part
        cur = head
        while right_head:
            tmp = right_head.next
            right_head.next = cur.next
            cur.next = right_head
            cur = right_head.next
            right_head = tmp


    def _reverse(self, head):
        if head is None or head.next is None:
            return head

        last = None
        cur = head
        while cur:
            nxt = cur.next
            cur.next = last
            last = cur
            cur = nxt

        return last


def print_list(head):
    while head:
        print head.val,
        head = head.next
    print "\n"


if __name__ == '__main__':
    s = Solution()
    head = ListNode(1)
    head.next = ListNode(2)
    head.next.next = ListNode(3)
    head.next.next.next = ListNode(4)
    #head.next.next.next.next = ListNode(5)

    print_list(head)

    s.reorderList(head)

    print_list(head)

