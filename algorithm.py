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

solution=Solution()
r1=solution.lengthOfLongestSubstring('abcabcdbtvtyyyyuuiutytb')
print(r1)