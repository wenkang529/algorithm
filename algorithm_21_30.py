import copy


# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def mergeTwoLists(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        if l1 is None:
            return l2
        if l2 is None:
            return l1

        if l1.val > l2.val:
            pre = r1 = l2
        else:
            pre = r1 = l1

        while l1 or l2:
            if l1 is None:
                result = l2
                pre.next = result
                return r1
            if l2 is None:
                result = l1
                pre.next = result
                return r1
            if l1 and l2:
                if l1.val <= l2.val:
                    result = l1
                    l1 = l1.next
                else:
                    result = l2
                    l2 = l2.next
                pre.next = result
                pre = result
        return r1

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
            self.generate(left - 1, right, result, string+'(')
        if left < right:
            self.generate(left, right - 1, result, string+')')

    def mergeKLists(self, lists):
        """
        :type lists: List[ListNode]
        :rtype: ListNode
        """
        result = []
        while True:
            flg = False
            for i in lists:
                if i != []:
                    flg = True
            if flg is False:
                break

            min = None
            for index, i in enumerate(lists):
                # print(i)
                if i == []:
                    continue
                if min is None:
                    min = i[0]
                    index_r = index
                else:
                    if min > i[0]:
                        min = i[0]
                        index_r = index
            if min is not None:
                lists[index_r].pop(0)
                result.append(min)
        return result

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

    def swapPairs(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """

        if head.next is None:
            return head
        p = head
        later = ListNode(0)
        later.next = head
        re = copy.deepcopy(later)

        while p:
            if p.next is None:
                later.next = p
                break
            b = p
            a = p.next
            new = p.next.next

            later.next = a
            a.next = b

            later = b
            p = new

        return re

    def swapPairs(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        root = ListNode(0)
        root.next = head
        last = root
        while last.next:
            if last.next.next:
                nex = last.next
                last.next, nex.next, last.next.next = nex.next, nex.next.next, nex  # exchange 3 node
                last = nex
            else:
                break
        return root.next

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
    def strStr(self, haystack, needle):
        """
        :type haystack: str
        :type needle: str
        :rtype: int
        """
        if needle=='' and haystack=='':
            return 0
        if needle not in haystack:
            return -1
        else:
            for i in range(len(haystack)):
                if haystack[i:i+len(needle)]==needle:
                    return i
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
        if (dividend==-2**31) and (divisor==-1):
            return (2**31-1)
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
    def findSubstring(self, s, words):
        """
        :type s: str
        :type words: List[str]
        :rtype: List[int]
        """
        res = []
        if not words or len(words[0]) == 0:
            return res
        wlen = len(words[0])
        window = wlen * len(words)
        for i in range(len(s) - window + 1):
            temp = words.copy()
            index = i
            while s[i:i + wlen] in temp:
                temp.remove(s[i:i + wlen])
                i += wlen
                if len(temp) == 0:
                    res.append(index)
        return res



a = ListNode(1)
b = ListNode(2)
c = ListNode(3)
d = ListNode(4)

a.next = b
b.next = c
c.next = d


solution = Solution()
result = solution.divide(-2147483648,-1)


print(result)
