# algorithm

## 11.container with most water

> Given *n* non-negative integers *a1*, *a2*, ..., *an*, where each represents a point at coordinate (*i*, *ai*). *n* vertical lines are drawn such that the two endpoints of line *i* is at (*i*, *ai*) and (*i*, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.
>
> Note: You may not slant the container and *n* is at least 2.



`my solution`

```python
# it erro
def maxArea(self, height):
  """
        :type height: List[int]
        :rtype: int
        """
  r_b=0
  r_e=len(height)-1

  r1=0

  while r_b<r_e:
    if height[r_b]<height[r_e]:
      if min(height[r_b],height[r_e])*(r_e-r_b)<min(height(r_b+1),height[r_e])*(r_e-r_b-1):
        r_b+=1
        else:
          if min(height[r_e],height[r_b])*(r_e-r_b)<min(height[r_e-1],height[r_b])*(r_e-r_b-1):
            r_e-=1
```

`attention`

- shold clear the key point: just need to move the shoter one,if move the longer one it must contain little water



`accept solution`

```python
def maxArea(self, height):
  """
        :type height: List[int]
        :rtype: int
        """
  n=len(height)
  b,e=0,n-1
  max_area=min(height[b],height[e])*(e-b)
  while b<e:
    if height[b]<height[e]:
      b+=1
      elif height[b]>=height[e]:
        e-=1
        max_area=max(max_area,min(height[e],height[b])*(e-b))
        return max_area
```

## 12.inter to roman

> Given an integer, convert it to a roman numeral.
>
> Input is guaranteed to be within the range from 1 to 3999.

`my solution`

```python
def intToRoman(self, num):
"""
:type num: int
:rtype: str
"""
M = ["", "M", "MM", "MMM"];
C = ["", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"]
X = ["", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"]
I = ["", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"]
return M[num // 1000] + C[(num % 1000) // 100] + X[(num % 100) // 10] + I[num % 10]
```

`attention`

- use table is more faster

`accept solution`

```
just like my solution
```

## 

## 13 roman to integer

> Given a roman numeral, convert it to an integer.
>
> Input is guaranteed to be within the range from 1 to 3999.

`my solution`

```python
def romanToInt(self, s):
  """
  :type s: str
  :rtype: int
  """
  M = ["", "M", "MM", "MMM"];
  C = ["", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"]
  X = ["", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"]
  I = ["", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"]
  if s=='':
  return 0
  r1=0
  j=0
  if s[0]=='M':
  while j<len(s) and s[j]=='M':
  j+=1
  if j>=len(s):
  m=s[0:]
  else:
  m=s[0:j]
  else:
  m=''
  i=j
  if j<len(s) and (s[j]=='C'or s[j]=='D'):
  while j<len(s) and (s[j]=='C' or s[j]=='D' or s[j]=='M'):
  j+=1
  if j>=len(s):
  c=s[i:]
  else:
  c=s[i:j]
  else:
  c=''
  i=j
  if j<len(s)and (s[j]=='X' or s[j]=='L'):
  while j<len(s) and (s[j]=='X' or s[j]=='L' or s[j]=='C'):
  j+=1
  if j>=len(s):
  x=s[i:]
  else:
  x=s[i:j]
  else:
  x=''
  i=j
  if j<len(s)and (s[j]=='I' or s[j]=='V'):
  ii=s[i:]
  else:
  ii=''
  return M.index(m)*1000+C.index(c)*100+X.index(x)*10+I.index(ii)
```

`attention`

- it's easy to ignore index out of range
- must consider ench step if index out of range
- can make code more short
- use `for i in range()`  or ` for i in s`  not have much diffrent

`accept solution`

```python
def romanToInt(self, s):
  """
        :type s: str
        :rtype: int
        """
  d = {'M': 1000, 'D': 500, 'C': 100, 'L': 50, 'X': 10, 'V': 5, 'I': 1}
  r1,p=0,'I'
  for i in range(len(s)-1,-1,-1):
    r1,p=r1+d[s[i]] if d[s[i]]>= d[p] else r1-d[s[i]],s[i]
    return r1
```

## longest common prefix

> Write a function to find the longest common prefix string amongst an array of strings.

`my solution`

```python
def longestCommonPrefix(self, strs):
  """
  :type strs: List[str]
  :rtype: str
  """
  if strs==[]:
  return ''
  cp=strs[0]
  pre=strs[0]
  for p in strs:
  i=0
  if p=='':
  return ''
  while i <len(p)and i<len(cp)and p[i]==cp[i]:
  i+=1
  cp=p[:i]
  pre=p
  return cp
```

`attention`

- consider if input is None
- use built-in-functions is more powerful (zip,set)

`accept solution`

```python
def longestCommonPrefix(self, strs):
    """
    :type strs: List[str]
    :rtype: str
    """
    if not strs:
        return ""

    for i, letter_group in enumerate(zip(*strs)):
        if len(set(letter_group)) > 1:
            return strs[0][:i]
    else:
        return min(strs
```

## 15.3Sum

> Given an array *S* of *n* integers, are there elements *a*, *b*, *c* in *S* such that *a* + *b* + *c* = 0? Find all unique triplets in the array which gives the sum of zero.
>
> **Note:** The solution set must not contain duplicate triplets.
>
> ```
> For example, given array S = [-1, 0, 1, 2, -1, -4],
>
> A solution set is:
> [
>   [-1, 0, 1],
>   [-1, -1, 2]
> ]
> ```

`my solution`

```python
def threeSum(self, nums):
    """
    :type nums: List[int]
    :rtype: List[List[int]]
    """
    length=len(nums)
    r1=[]
    for i in range(len(nums)):
        if i>=length-2:
            break
        for j in range(i+1,length):
            if j>=length-1:
                break
            for k in range(j+1,length):
                if k>=length:
                    break
                if nums[i]+nums[j]+nums[k]==0:
                    r2=[nums[i],nums[j],nums[k]]
                    if r1 !=[]:
                        flag=1
                        for pre in r1:
                            r3=set(r2)==set(pre)
                            r5=(pre[0]+pre[1]+pre[2]==r2[0]+r2[1]+r2[2])
                            if r3 and r5:
                                flag=0

                        if flag!=0:
                            r1.append(r2)
                    else:
                        r1.append(r2)
    return r1
```

`attention`

- my code time limited exceeded
- use much time to deal duplicate triplets
- it sort first,though sort may cost some time ,but it may be more easy

`accept solution`

```python
def threeSum(self, nums):
    res = []
    nums.sort()
    for i in xrange(len(nums)-2):
        if i > 0 and nums[i] == nums[i-1]:
            continue
        l, r = i+1, len(nums)-1
        while l < r:
            s = nums[i] + nums[l] + nums[r]
            if s < 0:
                l +=1 
            elif s > 0:
                r -= 1
            else:
                res.append((nums[i], nums[l], nums[r]))
                while l < r and nums[l] == nums[l+1]:
                    l += 1
                while l < r and nums[r] == nums[r-1]:
                    r -= 1
                l += 1; r -= 1
    return res
```

## 16.3sum closest

> Given an array *S* of *n* integers, find three integers in *S* such that the sum is closest to a given number, target. Return the sum of the three integers. You may assume that each input would have exactly one solution.
>
> ```
>     For example, given array S = {-1 2 1 -4}, and target = 1.
>
>     The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
> ```

`my solution`

```python
# Time Limit Exceeded
def threeSumClosest(self, nums, target):
  dis=None
  pre=None
  length=len(nums)
  for i in range(length):
    for j in range(i+1,length):
      for k in range(j+1,length):
        if k>=length:
          continue
          dis=abs(target-nums[i]-nums[j]-nums[k])
          if dis==None or pre==None or dis<pre:
            r1=nums[i]+nums[j]+nums[k]
            pre=dis
            return r1
```

`attention`

- use sort first
- first should add `r1=nums[0]+nums[1]+nums[2]`

`accept solution`

```python
    def threeSumClosest(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        nums.sort()
        r1=nums[0]+nums[1]+nums[2]
        length=len(nums)
        for i in range(length-2):
            j , k=i+1,length-1
            while j<k:
                sum=nums[i]+nums[j]+nums[k]
                if sum==target:
                    return sum
                if abs(sum-target)<abs(r1-target):
                    r1=sum
                if sum-target>0:
                    k-=1
                else:
                    j+=1
        return r1
```

# 17 letter combinations of a phone number

Given a digit string, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below.

![img](http://upload.wikimedia.org/wikipedia/commons/thumb/7/73/Telephone-keypad2.svg/200px-Telephone-keypad2.svg.png)

```
Input:Digit string "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

```

**Note:**
Although the above answer is in lexicographical order, your answer could be in any order you want.

`my solution`

```python
def letterCombinations(self, digits):
  r1=[]
  r2=[]
  mapping = {'2': 'abc', '3': 'def', '4': 'ghi', '5': 'jkl',
             '6': 'mno', '7': 'pqrs', '8': 'tuv', '9': 'wxyz'}
  for i in digits:
    c=mapping[i]
    r3=[]
    if r2==[]:
      for j in c:
        r3.append(j)
        else:
          for j in c:
            for k in r2:
              r3.append(k+j)
              r2=r3
              return r2
```

`attention`
- use dict more power
- use built-in function is more efficent
- ​


















