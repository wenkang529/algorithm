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

##22.



Given *n* pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given *n* = 3, a solution set is:

```
[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
```

## 

`my solution`



`attention`



`accept solution`

- two method

```python
    def generateParenthesis(self, n):
        """
        :type n: int
        :rtype: List[str]
        """
        res = []
        # number of valid left parenthesis
        l_valid = 1
        # number of remaining left parenthesis
        l = n - 1
        # number of remaining right parenthesis
        r = n
        # initial parenthesis string
        s = '('

        # recursive function for generating string s
        def generate(res, s, l_valid, l, r):
            # base case: if no remaining left parenthesis, add all remaining right parenthesis to s
            if l == 0:
                s += ')' * r
                res.append(s)
            # if have both remaining left parenthesis and valid left parenthesis
            # either add 1 left parenthesis to s or add 1 right parenthesis to s
            elif l != 0 and l_valid != 0:
                generate(res, s + '(', l_valid + 1, l - 1, r)
                generate(res, s + ')', l_valid - 1, l, r - 1)
            # if only have remaining left parenthesis not valid parenthesis
            # add 1 left parenthesis to s
            elif l != 0 and l_valid == 0:
                generate(res, s + '(', l_valid + 1, l - 1, r)

        generate(res, '(', l_valid, l, r)
        return res
    def generateParenthesis(self, n):
        """
        :type n: int
        :rtype: List[str]
        """
        if n == 0:
            return []
        left = right = n
        result = []
        self.generate(left, right, result, '')
        return result
    def generate(self, left, right, result, string):
        if left == 0 and right == 0:
            result.append(string)
            return
        if left:
            self.generate(left - 1, right , result, string+'(')
        if left < right:
            self.generate(left, right - 1, result, string+')')
```



## 23.Merge k Sorted Lists

Merge *k* sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

**Example:**

```
Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
```

`my solution`



`attention`

- so easy to use python sorted

`accept solution`

```python
    def mergeKLists(self, lists):
        """
        :type lists: List[ListNode]
        :rtype: ListNode
        """
        self.nodes = []
        head = point = ListNode(0)
        for l in lists:
            while l:
                self.nodes.append(l.val)
                l = l.next
        for x in sorted(self.nodes):
            point.next = ListNode(x)
            point = point.next
        return head.next
```









