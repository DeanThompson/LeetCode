#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
https://oj.leetcode.com/problems/remove-duplicates-from-sorted-list/
"""


# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    # @param head, a ListNode
    # @return a ListNode
    def deleteDuplicates(self, head):
        if head is None or head.next is None:
            return head

        pointer = head
        while pointer:
            while pointer.next and pointer.next.val == pointer.val:
                pointer.next = pointer.next.next

            pointer = pointer.next
        return head


def print_list(head):
    while head:
        print head.val,
        head = head.next
    print "\n"


if __name__ == '__main__':
    s = Solution()
    head = ListNode(1)
    head.next = ListNode(1)
    head.next.next = ListNode(2)
    head.next.next.next = ListNode(3)
    head.next.next.next.next = ListNode(3)

    print_list(head)

    s.deleteDuplicates(head)

    print_list(head)

