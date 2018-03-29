## 1.sum
> Given an array of integers,return indices of the two numbers such that they add up to a specific target.you may assume that each input would have exactly one solution.and you may not use the same elment twice.  
Example:   
Given two nums=[2,7,11,15] target=9
return [0,1]  
`my solution:`
```py
class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        a=[]
        for index,num in enumerate(nums):
            if num in a:
                return (a.index(num)),index
            a.append(target - num)
```
`attention`  

* requirement is return list
* use dict to seek is more power than use list
* add if...return in the begin of the code

`accept solution`

```py
class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        if len(nums)<=1:
            return False
        a={}
        for index,num in enumerate(nums):
            if num in a:
                return [a[num],index]
            a[target-num]=index
```

## 2.add two numbers
> you are given two non-empty linked lists representing two non-negative integers.the digits are stored in reverse order and each of their nodes contain a single digit. add the two numbers and return ir as a linked list. you may assume the two numbers do not contain any leading zero,expect the number 0 itself.  
`my solution`
```py
# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

######################################### my code ############################################# 

class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        begin_l1=l1
        begin_l2=l2
        result=ListNode(None)
        result1=result
        d=0
        while begin_l1 is not None:
            a=begin_l1.val
            if begin_l2==None:
                b=0
                begin_l2=None
            else:
                b = begin_l2.val
                begin_l2 = begin_l2.next
            c=a+b+d
            d=c//10
            y=c%10
            new=ListNode(y)
            result.next=new
            result=new
            begin_l1=begin_l1.next
        while begin_l2 is not None:
            b=begin_l2.val
            c=b+d
            d=c//10
            y=c%10
            new=ListNode(y)
            result.next=new
            result=new
            begin_l2=begin_l2.next
        if d  is not 0:
            new=ListNode(d)
            result.next=new
        return result1.next
```
`attention`

* care if carry(进位) is zero
* it asked the list must be the nodeList that they give.
* return the result1.next
* know clear how to edit the code structure(use `while l1 or l2 or carry`  not use  `while l1 is not None`)


`accept solution`

```py
def addTwoNumbers(self, l1, l2):
    dummy = cur = ListNode(0)
    carry = 0
    while l1 or l2 or carry:
        if l1:
            carry += l1.val
            l1 = l1.next
        if l2:
            carry += l2.val
            l2 = l2.next
        cur.next = ListNode(carry%10)
        cur = cur.next
        carry //= 10
    return dummy.next
```
## 3. Longest Substring Without Repeating Characters
> Given a string, find the length of the longest substring without repeating characters.  
Examples:  
Given "abcabcbb", the answer is "abc", which the length is 3.  
Given "bbbbb", the answer is "b", with the length of 1.  
Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.  

`my solution`

```py

def lengthOfLongestSubstring(self, s):
    """
    :type s: str
    :rtype: int
    """
    a=0
    dict={}
    length=0
    for index,i in enumerate(s):
        if i in dict:
            if dict[i]+1>a:
                a = dict[i]+1
        dict[i]=index
        if length < index-a+1:
            length=index-a+1
    return length
```
`attention`

* i foget to judge if a is the biggest one,so it's necessary to decide if the method is right and rigoros(严谨) 
* accept solution is more concise(简洁)

`accept solution`
```py
    def lengthOfLongestSubstring(self, s):
        start = maxLength = 0
        usedChar = {}
        
        for i in range(len(s)):
            if s[i] in usedChar and start <= usedChar[s[i]]:
                start = usedChar[s[i]] + 1
            else:
                maxLength = max(maxLength, i - start + 1)
            usedChar[s[i]] = i
        return maxLength
```
## 4.Median of Two Sorted Arrays
> There are two sorted arrays nums1 and nums2 of size m and n respectively.  
Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).  
Example 1:  
nums1 = [1, 3]  
nums2 = [2]   
The median is 2.0  
Example 2:  
nums1 = [1, 2]  
nums2 = [3, 4]  
The median is (2 + 3)/2 = 2.5


`my solution`
```py
    def findMedianSortedArrays(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        """
        m=len(nums1)
        n=len(nums2)
        i=0
        j=0
        num=0
        max=0
        min=0
        while True:
           if i<m and j<n:
               if nums1[i]<nums2[j]:
                   i+=1
                   num+=1
                   min=max
                   max=nums1[i-1]
               else:
                   j+=1
                   num+=1
                   min=max
                   max=nums2[j-1]
           elif i>=m:
                j += 1
                num += 1
                min = max
                max = nums2[j - 1]
           else:
               i += 1
               num += 1
               min = max
               max = nums1[i - 1]
           if num>=(m+n)//2+1:
               break

        if (m+n)%2==0:
           return (max+min)/2
        else:
           return max+0.0
```
`attention`

* forget to attention if the index out of limit
* it asked to return float but i didn't make diffrent from int to float `return (max+min)/2` or `return (max+min)/2.0`

`accept solution`  
[link](https://leetcode.com/problems/median-of-two-sorted-arrays/solution/)
```py
# 没有理解是什么意思

```
## 5. Longest Palindromic Substring
> Given a string s, find the longest palindromic(回文) substring in s. You may assume that the maximum length of s is 1000.  
Example:  
Input: "babad"  
Output: "bab"  
Note: "aba" is also a valid answer.  
Example:  
Input: "cbbd"  
Output: "bb"  

`my solution`
```py
    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        if len(s)<=1:
            return s
        r1=[]
        for i in range(len(s)):
                m=i-1
                j=i
                while m>=0 and j < len(s) and s[j]==s[m]:
                    j+=1
                    m-=1
                if j-m>2 and len(r1)< j-m-1:
                    r1=s[m+1:j]
                m=i-2
                j=i
                while m>=0 and j< len(s) and s[j]==s[m]:
                    j+=1
                    m-=1
                if j-m>3 and len(r1)< j-m-1:
                    r1=s[m+1:j]
        if r1==[]:
            r1=s[-1]
        return r1
```

`attention`

* be thoughtful,don't forget if the input is None or only one character
* my code run with more time,if change `while` to compare the two list will cost less time 

`accept`
```py
    def longestPalindrome(self, s):
        if len(s)==0:
        	return 0
        maxLen=1
        start=0
        for i in xrange(len(s)):
        	if i-maxLen >=1 and s[i-maxLen-1:i+1]==s[i-maxLen-1:i+1][::-1]:
        		start=i-maxLen-1
        		maxLen+=2
        		continue
        	if i-maxLen > = 0 and s[i-maxLen:i+1]==s[i-maxLen:i+1][::-1]:
        		start=i-maxLen
        		maxLen+=1
        return s[start:start+maxLen]
```













