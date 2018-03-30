## 21.merge two sorted lists

Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

**Example:**

```
Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
```

`my solution`

```python
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
```

`attention`

- creat one listnode first,so needn't compute which to return,and return .next

`accept solution`

```python
    def mergeTwoLists(self, l1, l2):

        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        dummy = ListNode(-1)
        cur = dummy
        while l1 and l2:
            if l1.val>l2.val:
                cur.next = l2
                l2 = l2.next
            else:
                cur.next = l1
                l1 = l1.next
            cur = cur.next
        if l1:
            cur.next = l1
        if l2:
            cur.next = l2
        return dummy.next
```

## 