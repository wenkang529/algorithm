# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def mergeTwoLists(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        if l1 is None:
            return l2
        if l2 is None:
            return l1

        if l1.val>l2.val:
            pre=r1=l2
        else:
            pre=r1=l1

        while l1 or l2:
            if l1 is None:
                result=l2
                pre.next=result
                return r1
            if l2 is None:
                result=l1
                pre.next=result
                return r1
            if l1 and l2:
                if l1.val<=l2.val:
                    result=l1
                    l1=l1.next
                else:
                    result=l2
                    l2=l2.next
                pre.next=result
                pre=result
        return r1
