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

## 24.Swap Nodes in Pairs

Given a linked list, swap every two adjacent nodes and return its head.

**Example:**

```
Given 1->2->3->4, you should return the list as 2->1->4->3.
```

**Note:**

- Your algorithm should use only constant extra space.
- You may **not** modify the values in the list's nodes, only nodes itself may be changed.

`accept solution`

```python
    def swapPairs(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        root=ListNode(0)
        root.next=head
        last=root
        while last.next:
            if last.next.next:
                nex = last.next
                last.next, nex.next, last.next.next = nex.next, nex.next.next, nex # exchange 3 node
                last = nex
            else:
                break
        return root.next
```

## 25.Reverse Nodes in k-Group

Given a linked list, reverse the nodes of a linked list *k* at a time and return its modified list.

*k* is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of *k* then left-out nodes in the end should remain as it is.


**Example:**

Given this linked list: `1->2->3->4->5`

For *k* = 2, you should return: `2->1->4->3->5`

For *k* = 3, you should return: `3->2->1->4->5`

**Note:**

- Only constant extra memory is allowed.
- You may not alter the values in the list's nodes, only nodes itself may be changed.

`accept solution`

```python
    def reverseKGroup(self, head, k):
        """
        :type head: ListNode
        :type k: int
        :rtype: ListNode
        """
        if k == 0 or k == 1: return head

        start_node = ListNode(0)
        start_node.next = tail = head
        reverse_list = [start_node]

        count = 0
        while tail != None:
            count += 1
            reverse_list.append(tail)
            next_node = tail.next
            if not count % k:
                for i in reversed(range(len(reverse_list))):
                    if i - 1 == 0:
                        reverse_list[i].next = next_node
                    else:
                        reverse_list[i].next = reverse_list[i - 1]
                reverse_list = [reverse_list[1]]
                count = 0

            tail = next_node

        return start_node.next
        
```

`attention`

- i have some trouble about how to deal with k,in accept solution use a list to save the node-list ,and  reversed it.
- just like the previous problem ,use the root.next as the first node.

## 26.Remove Duplicates from Sorted Array 

Given a sorted array *nums*, remove the duplicates [**in-place**](https://en.wikipedia.org/wiki/In-place_algorithm) such that each element appear only *once* and return the new length.

Do not allocate extra space for another array, you must do this by **modifying the input array in-place** with O(1) extra memory.

**Example 1:**

```
Given nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.

It doesn't matter what you leave beyond the returned length.
```

**Example 2:**

```
Given nums = [0,0,1,1,1,2,2,3,3,4],

Your function should return length = 5, with the first five elements of nums being modified to 0, 1, 2, 3, and 4 respectively.

It doesn't matter what values are set beyond the returned length.
```

**Clarification:**

Confused why the returned value is an integer but your answer is an array?

Note that the input array is passed in by **reference**, which means modification to the input array will be known to the caller as well.

Internally you can think of this:

```c++
// nums is passed in by reference. (i.e., without making a copy)
int len = removeDuplicates(nums);

// any modification to nums in your function would be known by the caller.
// using the length returned by your function, it prints the first len elements.
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
```

`my solution`

`accept solution`

```python
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if len(nums)<2:
            return len(nums)
        i=0
        j=1
        while j<len(nums):
            if nums[i]==nums[j]:
                j+=1
            else:
                nums[i+1]=nums[j]
                i+=1
                j+=1
        nums=nums[0:i+1]
        return i+1
```

`attention`

- it's important to deal how to use O(n) space
- use i,j  i is slow-runner,j is faster-runner. copy nums[j] to nums[i+1].



## 27.Remove Element

Given an array *nums* and a value *val*, remove all instances of that value [**in-place**](https://en.wikipedia.org/wiki/In-place_algorithm) and return the new length.

Do not allocate extra space for another array, you must do this by **modifying the input array in-place** with O(1) extra memory.

The order of elements can be changed. It doesn't matter what you leave beyond the new length.

**Example 1:**

```
Given nums = [3,2,2,3], val = 3,

Your function should return length = 2, with the first two elements of nums being 2.

It doesn't matter what you leave beyond the returned length.
```

**Example 2:**

```
Given nums = [0,1,2,2,3,0,4,2], val = 2,

Your function should return length = 5, with the first five elements of nums containing 0, 1, 3, 0, and 4.

Note that the order of those five elements can be arbitrary.

It doesn't matter what values are set beyond the returned length.
```

**Clarification:**

Confused why the returned value is an integer but your answer is an array?

Note that the input array is passed in by **reference**, which means modification to the input array will be known to the caller as well.

Internally you can think of this:

```cpp
// nums is passed in by reference. (i.e., without making a copy)
int len = removeElement(nums, val);

// any modification to nums in your function would be known by the caller.
// using the length returned by your function, it prints the first len elements.
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
```

`my solution`

```python
    def removeElement(self, nums, val):
        """
        :type nums: List[int]
        :type val: int
        :rtype: int
        """
        n=0
        for i in range(len(nums)-1,-1,-1):
            if nums[i]==val:
                nums.pop(i)
        return len(nums)
```

`attention`

- it's easy but if i want to through the array and delete remove element meantime.it will be wrong for index will skip the next element

## 28.Implement strStr()

Implement [strStr()](http://www.cplusplus.com/reference/cstring/strstr/).

Return the index of the first occurrence of needle in haystack, or **-1** if needle is not part of haystack.

**Example 1:**

```
Input: haystack = "hello", needle = "ll"
Output: 2
```

**Example 2:**

```
Input: haystack = "aaaaa", needle = "bba"
Output: -1
```

**Clarification:**

What should we return when `needle` is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when `needle` is an empty string. This is consistent to C's [strstr()](http://www.cplusplus.com/reference/cstring/strstr/) and Java's [indexOf()](https://docs.oracle.com/javase/7/docs/api/java/lang/String.html#indexOf(java.lang.String)).

`my solution`

```python
    def divide(self, dividend, divisor):
        """
        :type dividend: int
        :type divisor: int
        :rtype: int
        """
        if dividend==-2**23 and divisor==-1:
            return 2**23-1
        if divisor>0 and dividend>0:
            s=1
        elif divisor<0 and dividend<0:
            s=1
        else:
            s=-1
        m=abs(dividend)
        n=abs(divisor)
        r=0
        while m-n>=0:
            r+=1
            m=m-n
        if s==-1:
            return 0-r
        else:
            return r
```

`attention`

- my solution run out of time
- use  move left and the value bigger twice , it 's more powerful

`accept solution`

```python
    def divide(self, dividend, divisor):
        """
        :type dividend: int
        :type divisor: int
        :rtype: int
        """
        if (dividend<0) is (divisor<0):
            s=1
        else:
            s=-1
        if dividend==-2**31 and divisor==-1:
            return 2**31-1
        divisor=abs(divisor)
        dividend=abs(dividend)

        ys=dividend
        num=0
        if dividend<divisor:
            return 0
        while ys>=divisor:
            ys,n=self.divide_f(ys,divisor)
            num=num+n
        if s==-1:
            num=0-num
        return num

    def divide_f(self,dividend,divisor):
        new=divisor
        n=1
        while new<dividend:
            new<<=1
            if new<=dividend:
                n+=n
            else:
                new>>=1
                break
        ys=dividend-new
        return ys,n

```

## 30.Substring with Concatenation of All Words

You are given a string, **s**, and a list of words, **words**, that are all of the same length. Find all starting indices of substring(s) in **s** that is a concatenation of each word in **words** exactly once and without any intervening characters.

**Example 1:**

```
Input:
  s = "barfoothefoobarman",
  words = ["foo","bar"]
Output: [0,9]
Explanation: Substrings starting at index 0 and 9 are "barfoor" and "foobar" respectively.
The output order does not matter, returning [9,0] is fine too.
```

**Example 2:**

```
Input:
  s = "wordgoodstudentgoodword",
  words = ["word","student"]
Output: []
```

`my solution`

```python
    def findSubstring(self, s, words):
        """                                          
        :type s: str                                 
        :type words: List[str]                       
        :rtype: List[int]                            
        """                                          
        res=[]                                       
        if not words or len(words[0])==0:            
            return  res                              
        wlen=len(words[0])                                                           
        window=wlen*len(words)                             
        for i in range(len(s)-window+1):                  
            temp=words.copy()                              
            index=i                                  
            while s[i:i+wlen] in  temp:             
                temp.remove(s[i:i+wlen])             
                i+=wlen                              
                if len(temp)==0:                     
                    res.append(index)                
        return res  
```

`attention`

- here is a important problem :how to judge two array ,they are equal but have different order.
- i don't have a good idea,one poor solution ,compare one element whether in another list and remove the element from the second list.

