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
    def mergeKLists(self, lists):
        """
        :type lists: List[ListNode]
        :rtype: ListNode
        """
        result=[]
        while True:
            flg=False
            for i in lists:
                if i != []:
                    flg=True
            if flg is False:
                break

            min=None
            for index,i in enumerate(lists):
                # print(i)
                if i == []:
                    continue
                if min is None:
                    min=i[0]
                    index_r=index
                else:
                    if min>i[0]:
                        min=i[0]
                        index_r=index
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
        p=head
        later=ListNode(0)
        later.next=head
        re=copy.deepcopy(later)

        while p:
            if p.next is None:
                later.next=p
                break
            b=p
            a=p.next
            new=p.next.next

            later.next=a
            a.next=b

            later=b
            p=new


        return re









a=ListNode(1)
b=ListNode(2)
c=ListNode(3)
d=ListNode(4)

a.next=b
b.next=c
c.next=d







solution=Solution()
result=solution.swapPairs(a)

while result.next:
    print(result.val)
    result=result.next
print(result.val)