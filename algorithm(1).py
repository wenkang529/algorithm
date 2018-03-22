
class Solution(object):
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

solution=Solution()
r1=solution.threeSum([5,-9,-11,9,9,-4,14,10,-11,1,-13,11,10,14,-3,-3,-4,6,-15,6,6,-13,7,-11,-15,10,-8,13,-14,-12,12,6,-6,8,0,10,-11,-8,-2,-6,8,0,12,3,-9,-6,8,3,-15,0,-6,-1,3,9,-5,-5,4,2,-15,-3,5,13,-11,7,6,-4,2,11,-5,7,12,-11,-15,1,-1,-9,10,-8,1,2,8,11,-14,-4,-3,-12,-2,8,5,-1,-9,-4,-3,-13,-12,-12,-10,-3,6,1,12,3,-3,12,11,11,10])
print(r1)