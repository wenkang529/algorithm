import time

class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        if len(nums)<=1:
            return False
        a={}
        for index,num in enumerate(nums):
            if num in a:
                return [a[num],index]
            a[target-num]=index
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

        # u=0
        # r=[]
        # result=[]
        # for index,i in enumerate(l1):
        #     if index>len(l2)-1:
        #         r.append((i+u)%10)
        #         u=(i+u)//10
        #     else:
        #         r1=i+l2[index]+u
        #         r.append(r1%10)
        #         u=r1//10
        # for i in range(len(r)):
        #     result.append(r[len(r)-1-i])
        # return result
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
            if length<index-a+1:
                length=index-a+1
        return length
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
                while m>=0 and j<len(s) and s[j]==s[m]:
                    j+=1
                    m-=1
                if j-m>2 and len(r1)<j-m-1:
                    r1=s[m+1:j]
                m=i-2
                j=i
                while m>=0 and j<len(s) and s[j]==s[m]:
                    j+=1
                    m-=1
                if j-m>3 and len(r1)<j-m-1:
                    r1=s[m+1:j]
        if r1==[]:
            r1=s[-1]
        return r1
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
                while i+j<len(s):
                    r1+=s[i+j]
                    if i+j+(numRows-1-i)*2<len(s):
                        r1+=s[i+j+(numRows-1-i)*2]
                    j+=skip
        return r1
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
        if x<0:
            flag=-1
            x=x*-1
        else:
            flag=1

        while x!=0:
            r2=r2*10+x%10
            x=x//10
        return r2*flag*(r2<2**31)
    def myAtoi(self, str):
        """
        :type str: str
        :rtype: int
        """
        j = 0
        r1 = 0
        flag = 1
        if len(str) == 0:
            return 0
        while j < len(str) and str[j] == ' ':
            j += 1
        if j < len(str) and str[j] in '+-':
            flag = 1 - 2 * (str[j] == '-')
            j += 1
        if j < len(str):
            for i in str[j:]:
                if i.isdigit():
                    r1 = r1 * 10 + int(i)
                    j += 1
                else:
                    break
        if r1 * flag > 2 ** 31 - 1:
            return 2 ** 31 - 1
        if r1 * flag < -(2 ** 31):
            return -(2 ** 31)
        return r1 * flag
    def isPalindrome(self, x):
        """
        :type x: int
        :rtype: bool
        """
        if x<0 or (x!=0 and x%10==0):
            return False
        sum=0
        while sum<x:
            sum=sum*10+x%10
            x=x//10
        return (sum==x) or (x==sum//10)




solution=Solution()
r1=solution.isPalindrome(12021)
print(r1)