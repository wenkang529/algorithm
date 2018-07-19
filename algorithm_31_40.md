## 31.Next Permutation

Implement **next permutation**, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

The replacement must be **in-place** and use only constant extra memory.

Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.

`1,2,3` → `1,3,2`
`3,2,1` → `1,2,3`
`1,1,5` → `1,5,1`

`my solution`

```python

```

`attention`

- didn't get what mean of this problem

`accept solution`

```python

```

## 32.Longest Valid Parentheses

Given a string containing just the characters `'('` and `')'`, find the length of the longest valid (well-formed) parentheses substring.

**Example 1:**

```
Input: "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()"
```

**Example 2:**

```
Input: ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()"
```

`my solution`

```python

```

`attention`

- dynamic programming

## 33.Search in Rotated Sorted Array

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., `[0,1,2,4,5,6,7]` might become `[4,5,6,7,0,1,2]`).

You are given a target value to search. If found in the array return its index, otherwise return `-1`.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of *O*(log *n*).

**Example 1:**

```
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
```

**Example 2:**

```
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
```

`my solution`

```python
    def search(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        low,high=0,len(nums)-1

        while low<=high:
            mid = (low + high) // 2
            if nums[mid]==target:
                return mid

            if nums[mid]>=nums[low]:
                if nums[low]<=target<=nums[mid]:
                    high=mid-1
                else:
                    low=mid+1
            else:
                if nums[high]>=target>=nums[mid]:
                    low=mid+1
                else:
                    high=mid-1
        return -1
```

`attention`

- binary search solution


## 34.Find First and Last Position of Element in Sorted Array

------

Given an array of integers `nums` sorted in ascending order, find the starting and ending position of a given `target` value.

Your algorithm's runtime complexity must be in the order of *O*(log *n*).

If the target is not found in the array, return `[-1, -1]`.

**Example 1:**

```
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
```

**Example 2:**

```
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
```

`my solution`

```python
    def searchRange(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        if len(nums)<1:
            return [-1,-1]
        if target<nums[0] or target>nums[-1]:
            return [-1,-1]
        low,high=0,len(nums)-1
        r1=-1

        if target==nums[-1]:
            r1=len(nums)-1
        else:
            while low<high:
                mid=(low+high)//2
                if nums[mid]==target:
                    r1=mid
                    break
                if mid==low:
                    break
                if nums[mid]<target:
                    low=mid
                else:
                    high=mid

        if r1==-1:
            return [r1,r1]
        else:
            begin=r1-1
            end=r1+1
            while True:
                if begin<0 or nums[begin]!=target:
                    rb=begin+1
                    break
                else:
                    begin-=1
            while True:
                if end>=len(nums) or nums[end]!=target:
                    re=end-1
                    break
                else:
                    end+=1
        return [rb,re]
        
```

`attention`

- binary solution
- judge the edge of array ,if code enter the loop forever

```

```

## 35.Search Insert Position

Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

**Example 1:**

```
Input: [1,3,5,6], 5
Output: 2
```

**Example 2:**

```
Input: [1,3,5,6], 2
Output: 1
```

**Example 3:**

```
Input: [1,3,5,6], 7
Output: 4
```

**Example 4:**

```
Input: [1,3,5,6], 0
Output: 0
```

```python
    def searchInsert(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        if target<=nums[0]:
            return 0
        if target>nums[-1]:
            return len(nums)
        if target==nums[-1]:
            return len(nums)-1
        low,high=0,len(nums)-1
        while low<=high:
            mid=(low+high)//2
            if nums[mid]==target:
                return mid
            if low==high-1:
                if (nums[high]==target)  or (nums[high]>target>nums[low]):
                    return high
                if nums[low]==target:
                    return low
            if nums[mid]<target:
                low=mid
            else:
                high=mid

```

`attention`

- binary solution

## 36.Valid Sudoku

Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated **according to the following rules**:

1. Each row must contain the digits `1-9` without repetition.
2. Each column must contain the digits `1-9` without repetition.
3. Each of the 9 `3x3` sub-boxes of the grid must contain the digits `1-9` without repetition.

![img](https://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/Sudoku-by-L2G-20050714.svg/250px-Sudoku-by-L2G-20050714.svg.png)
A partially filled sudoku which is valid.

The Sudoku board could be partially filled, where empty cells are filled with the character `'.'`.

**Example 1:**

```
Input:
[
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
Output: true
```

**Example 2:**

```
Input:
[
  ["8","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being 
    modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
```

**Note:**

- A Sudoku board (partially filled) could be valid but is not necessarily solvable.
- Only the filled cells need to be validated according to the mentioned rules.
- The given board contain only digits `1-9` and the character `'.'`.
- The given board size is always `9x9`.

`my solution`

```python
    def isValidSudoku(self, board):
        """
        :type board: List[List[str]]
        :rtype: bool
        """
        for i in range(9):
            tem1 = [0] * 10
            tem2 = [0] * 10
            for j in range(9):
                # print('lie',j,i,board[j][i])
                if board[j][i]!='.':
                    if tem1[int((board[j][i]))]==1:
                        return False
                    else:
                        tem1[int(board[j][i])]=1
                if board[i][j]!='.':
                    # print('hang:',i,j,board[i][j])
                    if tem2[int(board[i][j])]==1:
                        return False
                    else:
                        tem2[int(board[i][j])]=1
        # print('hang right')
        for i in range(3):
            for j in range(3):
                tem=[0]*10
                for k in range(3):
                    for l in range(3):
                        if board[k+i*3][l+j*3]!='.':
                            if tem[int(board[k+i*3][l+j*3])]==1:
                                return False
                            else:
                                tem[int(board[k+i*3][l+j*3])]=1
        return True
```

`attention`

- easy  for row and column can be in one loop

## 37.Sudoku Solver

Write a program to solve a Sudoku puzzle by filling the empty cells.

A sudoku solution must satisfy **all of the following rules**:

1. Each of the digits `1-9` must occur exactly once in each row.
2. Each of the digits `1-9` must occur exactly once in each column.
3. Each of the the digits `1-9` must occur exactly once in each of the 9 `3x3` sub-boxes of the grid.

Empty cells are indicated by the character `'.'`.

![img](https://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/Sudoku-by-L2G-20050714.svg/250px-Sudoku-by-L2G-20050714.svg.png)
A sudoku puzzle...

![img](https://upload.wikimedia.org/wikipedia/commons/thumb/3/31/Sudoku-by-L2G-20050714_solution.svg/250px-Sudoku-by-L2G-20050714_solution.svg.png)
...and its solution numbers marked in red.

**Note:**

- The given board contain only digits `1-9` and the character `'.'`.
- You may assume that the given Sudoku puzzle will have a single unique solution.
- The given board size is always `9x9`.

`my solution`

```python
    def solveSudoku(self, board):
        """
        :type board: List[List[str]]
        :rtype: void Do not return anything, modify board in-place instead.
        """
        unfilled = []
        for i in range(len(board)):
            for j in range(len(board[0])):
                if board[i][j] == '.':
                    unfilled.append((i, j))
        found = [False]
        self.dfs(board, unfilled, 0, found)
        return
        
        
        
    def dfs(self, board, unfilled, i, found):
        if i == len(unfilled):
            found[0] = True
            return
        s, t = unfilled[i]
        for k in range(1, 10, 1):
            board[s][t] = str(k)
            if self.isValidAdd(board, s, t):
            #if self.isValidSudoku(board):
                self.dfs(board, unfilled, i + 1, found)
            if found[0] == False:
                board[s][t] = '.'
            else:
                break
        return
```

`attention`

- DFS

## 38.Count and Say

The count-and-say sequence is the sequence of integers with the first five terms as following:

```
1.     1
2.     11
3.     21
4.     1211
5.     111221
```

`1` is read off as `"one 1"` or `11`.
`11` is read off as `"two 1s"` or `21`.
`21` is read off as `"one 2`, then `one 1"` or `1211`.

Given an integer *n*, generate the *n*th term of the count-and-say sequence.

Note: Each term of the sequence of integers will be represented as a string.

**Example 1:**

```
Input: 1
Output: "1"
```

**Example 2:**

```
Input: 4
Output: "1211"
```

`my solution`

```python
    def countAndSay(self, n):
        """
        :type n: int
        :rtype: str
        """
        first='1'
        if n==1:
            return '1'
        for m in range(n-1):
            count=0
            bef=first[0]
            r1=''
            for j in first:
                if j==bef:
                    count+=1
                else:
                    r1=r1+str(count)+str(bef)
                    bef=j
                    count=1
            if count!=0:
                r1 = r1 + str(count) + str(bef)
            first=r1
        return r1
```

`attention`

- easy

## 39.Combination Sum

Given a **set** of candidate numbers (`candidates`) **(without duplicates)** and a target number (`target`), find all unique combinations in `candidates` where the candidate numbers sums to `target`.

The **same** repeated number may be chosen from `candidates` unlimited number of times.

**Note:**

- All numbers (including `target`) will be positive integers.
- The solution set must not contain duplicate combinations.

**Example 1:**

```
Input: candidates = [2,3,6,7], target = 7,
A solution set is:
[
  [7],
  [2,2,3]
]
```

**Example 2:**

```
Input: candidates = [2,3,5], target = 8,
A solution set is:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
```

`my solution`

```python

```

`attention`

- backtracking

`accept solution`

```python
    def combinationSum(self, candidates, target):
        ans = []
        candidates.sort()
        def backTracking(candidates, target, res):
            if sum(res) == target:
                ans.append(res)
                return

            if sum(res) > target:
                return

            for index, val in enumerate(candidates):
                if sum(res) + val > target: break
                backTracking(candidates[index:], target, res + [val])

        backTracking(candidates, target, [])
        return ans
```

## 40.Combination Sum II

Given a collection of candidate numbers (`candidates`) and a target number (`target`), find all unique combinations in `candidates` where the candidate numbers sums to `target`.

Each number in `candidates` may only be used **once** in the combination.

**Note:**

- All numbers (including `target`) will be positive integers.
- The solution set must not contain duplicate combinations.

**Example 1:**

```
Input: candidates = [10,1,2,7,6,1,5], target = 8,
A solution set is:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
```

**Example 2:**

```
Input: candidates = [2,5,2,1,2], target = 5,
A solution set is:
[
  [1,2,2],
  [5]
]
```

`my solution`

```python

```

`attention`

- same to 39.
- use DFS

`accept solution`

```python
    def combinationSum2(self, candidates, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        candidates.sort()
        stack=[[[],0,0]]
        res=[]
        visited={}
        
        while(stack):
            lst,covered,index=stack.pop()
            
            if covered==target:
                res.append(lst)
                continue
                
            for i in range(index,len(candidates)):
                if covered+candidates[i]<=target and tuple(lst+[candidates[i]]) not in visited:
                    stack.append([lst+[candidates[i]],covered+candidates[i],i+1])
                    visited[tuple(lst+[candidates[i]])]=1

        
        return res
```

