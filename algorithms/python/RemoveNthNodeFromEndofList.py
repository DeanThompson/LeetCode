#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
https://oj.leetcode.com/problems/remove-nth-node-from-end-of-list/
"""


# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    # @return a ListNode
    def removeNthFromEnd(self, head, n):
        if head is None:
            return None

        dummy = ListNode(0)
        dummy.next = head
        slow = dummy
        fast = dummy

        # fast be n nodes ahead than slow
        while n > 0:
            fast = fast.next
            n -= 1
     
        # move fast and slow in the same pace
        # slow will be (n-1)th node from the end while fast reach the end
        while fast and fast.next:
            fast = fast.next
            slow = slow.next

        # remove slow.next
        slow.next = slow.next.next

        return dummy.next

    def removeNthFromEnd2(self, head, n):
        if head is None:
            return None

        fast = head
        while n > 0:
            fast = fast.next
            n -= 1

        if fast is None:
            return head.next

        slow = head
        while fast and fast.next:
            fast = fast.next
            slow = slow.next
        slow.next = slow.next.next
        return head


def _print_list(head):
    while head:
        print head.val, 
        head = head.next

    print "\n"


if __name__ == '__main__':
    head = ListNode(1)
    node = head
    for i in [2, 3, 4, 5, 6, 7, 8]:
        node.next = ListNode(i)
        node = node.next

    _print_list(head)

    s = Solution()
    #head = s.removeNthFromEnd(head, 7)
    head = s.removeNthFromEnd2(head, 8)

    _print_list(head)

