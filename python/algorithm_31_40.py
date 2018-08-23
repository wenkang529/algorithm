class Solution(object):
    def nextPermutation(self, nums):
        perm, l = False, len(nums) - 2
        while 0 <= l:
            r = len(nums) - 1
            while l < r and nums[r] <= nums[l]: r -= 1
            if r <= l: l -= 1
            else: nums[l], nums[l + 1:], perm = nums[r], sorted(nums[l + 1:r] + [nums[l]] + nums[r + 1:]), True; break
        if not perm: nums.sort()
    def longestValidParentheses(self, s):
        """
        :type s: str
        :rtype: int
        """
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
                # if self.isValidSudoku(board):
                self.dfs(board, unfilled, i + 1, found)
            if found[0] == False:
                board[s][t] = '.'
            else:
                break
        return
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
    def combinationSum(self, candidates, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
    def combinationSum2(self, candidates, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """















solution=Solution()
result=solution.countAndSay(5)
print(result)