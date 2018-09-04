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
```c++
# DP算法
   //找出两个有序数组的中位数
    //中位数有两中,当为偶数时,中位数可以是两个数的平均,也可以分为上中位数和下中位数
    //一个比较好的方法是转化为求递i个小的数
    //getkth()中的k 范围是1:n  不能是0 否则会出现找不到文件malloc.c
    double find_mid_num(vector<int> arr1, vector<int> arr2)
    {
        int len1 = arr1.size();
        int len2 = arr2.size();
        //因为要求求中位数,所以如果是奇数那么这两个是一样的
        int left = (len1 + len2 + 1) / 2;
        int right = (len1 + len2 + 2) / 2;
        return (getkth(arr1, 0, len1 - 1, arr2, 0, len2 - 1, left) + getkth(arr1, 0, len1 - 1, arr2, 0, len2 - 1, right)) * 0.5;
    }
    int getkth(vector<int> arr1, int start1, int end1, vector<int> arr2, int start2, int end2, int k)
    {
        int len1 = end1 - start1 + 1;
        int len2 = end2 - start2 + 1;
        if (len1 > len2)
        {
            return getkth(arr2, start2, end2, arr1, start1, end1, k);
        }
        if (len1 == 0)
        {
            return arr2[start2 + k - 1];
        }
        if (k == 1)
        {
            return arr1[start1] < arr2[start2] ? arr1[start1] : arr2[start2];
        }
        //i,j 为切到i,j的位置
        int i = start1 - 1 + (len1 < (k / 2) ? len1 : (k / 2));
        int j = start2 - 1 + (len2 < (k / 2) ? len2 : (k / 2));
        if (arr1[i] > arr2[j])
        {
            return getkth(arr1, start1, end1, arr2, j + 1, end2, k - (j - start2 + 1));
        }
        else
        {
            return getkth(arr1, i + 1, end1, arr2, start2, end2, k - (i - start1 + 1));
        }
    }

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
string longestPalindrome(string s)
{
    int len = s.size();
    if (len < 2)
        return s;
    bool d[len][len];
    //init
    for(int i=0;i<len;i++){
        d[i][i]=true;
        if(i<len-1)
            d[i][i+1]=s[i]==s[i+1]?true:false;
    }
    //dp
    for(int i=len-3;i>=0;i--){
        for(int j=i+2;j<len;j++){
            d[i][j]=(s[i]==s[j]&&d[i+1][j-1]);
        }
    }
    //return the longest palindrome
    int maxnum=0;
    int m=0,n=0;
    for (int i=0;i<len;i++){
        for(int j=i;j<len;j++){
            if(d[i][j]&&j-i+1>maxnum){
                maxnum=j-i+1;
                m=i;
                n=j;
            }
        }
    }
    // cout<< maxnum<<endl;
    return s.substr(m,n-m+1);  //string.substr(begin,size)
}
```
# 6. ZigZag Conversion

> The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)  
```py
'''
P   A   H   N  
A P L S I I G  
Y   I   R 
'''
```
And then read line by line: "PAHNAPLSIIGYIR"  
Write the code that will take a string and make this conversion given a number of rows:  
string convert(string text, int nRows);  
convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".

`my solution`
```py
    def convert(self, s, numRows):
        """
        :type s: str
        :type numRows: int
        :rtype: str
        """
        if numRows==1:
            return s
        r1=''
        skip=2*numRows-2
        for i in range(numRows):
            if i==0:
                j = 0
                while i + j < len(s):
                    r1+=s[i + j]
                    j += skip
            if i==numRows-1:
                j = 0
                while i + j < len(s):
                    r1+=s[i + j]
                    j += skip
            if i !=0 and i!=numRows-1 :
                j=0
                while i+j< len(s):
                    r1+=s[i+j]
                    if i+j+(numRows-1-i)*2< len(s):
                        r1+=s[i+j+(numRows-1-i)*2]
                    j+=skip
        return r1
```

`attention`

* The most important is to care any impossible of input (I have forget to thought about the input is only one character)
* 
`accept solution`
```py

```
# 7.Reverse integer

> Given a 32-bit signed integer, reverse digits of an integer.  
Example 1:  
Input: 123  
Output:  321  
Example 2:  
Input: -123  
Output: -321  
Example 3:  
Input: 120  
Output: 21  
Note:  
Assume we are dealing with an environment which could only hold integers within the 32-bit signed integer range. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

`my solution`
```py
def reverse(self, x):
    """
    :type x: int
    :rtype: int
    """
    a=10
    r1=[]
    r2=0
    if x>2**31:
        return 0
    if x<0:
        flag=-1
        x=x*-1
    else:
        flag=1

    while x!=0:
        r1.append(x%10)
        x=x//10
    length=len(r1)
    for i in r1:
        r2+=i*10**(length-1)
        length-=1
    if r2>2**31:
        return 0
    return r2*flag
```
`attention`
* use r2*10+r1%10 is more efficent and concise(简洁) 
* use (r2<2**31) efficent
* don't forget to deal with the input/output overflows(溢出)
`accept solution`
```py
def reverse(self,x):
    '''
    :param x:int
    :return: int
    '''
    a=10
    r1=[]
    r2=0
    if x>2**31:
        return 0
    if x< 0:
        flag=-1
        x=x*-1
    else:
        flag=1
    while x!=0:
        r2=r2*10+x%10
        x=x//10
    return r2*flag*(r2< 2**31)
```
## 8.String to Integer(stoi)

> Implement atoi to convert a string to an integer.  
Hint: Carefully consider all possible input cases. If you want a challenge, please do not see below and ask yourself what are the possible input cases.  
Notes: It is intended for this problem to be specified vaguely (ie, no given input specs). You are responsible to gather all the input requirements up front.   
Requirements for atoi:  
The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.  
The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.  
If the first sequence of non-whitespace characters in str is not a valid   integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.   
If no valid conversion could be performed, a zero value is returned. If the correct value is out of the range of representable values, INT_MAX (2147483647) or INT_MIN (-2147483648) is returned.

` my solution`
```py
    def myAtoi(self, str):
        """
        :type str: str
        :rtype: int
        """
        j=0
        r1=0
        flag=1
        if len(str)==0:
            return 0
        while j< len(str) and str[j]==' ' :
            j+=1
        if j< len(str)and str[j]   in '+-':
            flag=1-2*(str[j]=='-')
            j+=1
        if j< len(str):
            for i in str[j:]:
                if i.isdigit():
                    r1=r1*10+int(i)
                    j+=1
                else:
                    break
   
        if r1*flag>2**31-1:
            return 2**31-1
        if r1*flag<-(2**31):
            return -(2**31)
        return r1*flag
```
`attention`
* the most important to care about is to thoughtful any kind of imput


`accept solution`
```

```

## 9.Palindrome Number

> Determine whether an integer is a palindrome. Do this without extra space.  
Some hints:  
Could negative integers be palindromes? (ie, -1)  
If you are thinking of converting the integer to string, note the restriction of using extra space.  
You could also try reversing an integer. However, if you have solved the problem "Reverse Integer", you know that the reversed integer might overflow. How would you handle such case?  
There is a more generic way of solving this problem.

`solution`
```
    def isPalindrome(self, x):
        """
        :type x: int
        :rtype: bool
        """
        if x< 0 or (x!=0 and x%10==0):
            return False
        sum=0
        while sum< x:
            sum=sum*10+x%10
            x=x//10
        return (sum==x) or (x==sum//10)

```
`attention`
```
* use less space
```

## 9 Regular Expression Matching (DP:dynamic programming)

Given an input string (`s`) and a pattern (`p`), implement regular expression matching with support for `'.'` and `'*'`.

```
'.' Matches any single character.
'*' Matches zero or more of the preceding element.
```

The matching should cover the **entire** input string (not partial).

**Note:**

- `s` could be empty and contains only lowercase letters `a-z`.
- `p` could be empty and contains only lowercase letters `a-z`, and characters like `.` or `*`.

**Example 1:**

```
Input:
s = "aa"
p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
```

**Example 2:**

```
Input:
s = "aa"
p = "a*"
Output: true
Explanation: '*' means zero or more of the precedeng element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
```

**Example 3:**

```
Input:
s = "ab"
p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
```

**Example 4:**

```
Input:
s = "aab"
p = "c*a*b"
Output: true
Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore it matches "aab".
```

**Example 5:**

```
Input:
s = "mississippi"
p = "mis*is*p*."
Output: false
```

`my solution`

```

```
`attention`

[dynamic programming](http://www.cnblogs.com/steven_oyj/archive/2010/05/22/1741374.html)  
[pic to understand DP](http://blog.csdn.net/mrlevo520/article/details/75676160)

write table and find the relationship

`accept solution`
```
def isMatch(self, s, p):
    # The DP table and the string s and p use the same indexes i and j, but
    # table[i][j] means the match status between p[:i] and s[:j], i.e.
    # table[0][0] means the match status of two empty strings, and
    # table[1][1] means the match status of p[0] and s[0]. Therefore, when
    # refering to the i-th and the j-th characters of p and s for updating
    # table[i][j], we use p[i - 1] and s[j - 1].

    # Initialize the table with False. The first row is satisfied.
    table = [[False] * (len(s) + 1) for _ in range(len(p) + 1)]

    # Update the corner case of matching two empty strings.
    table[0][0] = True

    # Update the corner case of when s is an empty string but p is not.
    # Since each '*' can eliminate the charter before it, the table is
    # vertically updated by the one before previous. [test_symbol_0]
    for i in range(2, len(p) + 1):
        table[i][0] = table[i - 2][0] and p[i - 1] == '*'

    for i in range(1, len(p) + 1):
        for j in range(1, len(s) + 1):
            #deal without '*'
            if p[i - 1] != "*":
                # Update the table by referring the diagonal element.
                table[i][j] = table[i - 1][j - 1] and \
                              (p[i - 1] == s[j - 1] or p[i - 1] == '.')
            else:
                # Eliminations (referring to the vertical element)
                # Either refer to the one before previous or the previous.
                # I.e. * eliminate the previous or count the previous.
                # [test_symbol_1]
                table[i][j] = table[i - 2][j] or table[i - 1][j]

                # Propagations (referring to the horizontal element)
                # If p's previous one is equal to the current s, with
                # helps of *, the status can be propagated from the left.
                # [test_symbol_2]
                if p[i - 2] == s[j - 1] or p[i - 2] == '.':
                    table[i][j] |= table[i][j - 1]

    return table[-1][-1]
```
