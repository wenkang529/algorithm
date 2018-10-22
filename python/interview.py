
if __name__ == "__main__":
    # 读取第一行的n
    n,m = input().split()
    n=int(n)
    m=int(m)
    arr1=list(input().split())
    arr1=map(int,arr1)
    arr1=list(arr1)

    arr2=list(input().split())
    arr2=map(int,arr2)
    arr2=list(arr2)

    mydic={}
    for i in range(n):
        mydic[arr1[i]]=arr2[i]
    
    arr1.sort(reverse=1)

    num=0
    i=0
    result=0
    while(num<m):
        if(mydic[arr1[i]]>0):
            num+=arr1[i]
            mydic[arr1[i]]-=1
            result+=1
        elif i<n:
            i+=1
        else:
            result=sum(arr2)
            break
    print(result)

# def quicksort(nums):
#     if len(nums) <= 1:
#         return nums
#     less = []
#     greater = []
#     base = nums.pop()
#     for x in nums:
#         if x < base:
#             less.append(x)
#         else:
#             greater.append(x)
#     return quicksort(less) + [base] + quicksort(greater)




# if __name__ == "__main__":
#     # 读取第一行的n
#     n= input().split()
#     arr1=list(input().split())
#     print(arr1)
#     print(arr1.sort())

    


N = int(input())

nums = list(map(str,input().split()))

for i in range(len(nums)):
    for j in range(1,len(nums)-i):
        if nums[j-1]+nums[j] < nums[j]+nums[j-1]:
            nums[j-1],nums[j] = nums[j],nums[j-1]
res = int("".join(nums))
print(res)


# import datetime
# N = int(input())
# ll=[]
# for i in range(N):
#     arr1=list(input().split())
#     arr1=map(int,arr1)
#     arr1=list(arr1)
#     ll.append(arr1)

# for i in ll:
#     if (i[1]>i[3]) or (i[1]==i[3] and i[2]>i[4]):
#         year=i[0]+1
#     else:
#         year=i[0]

#     d1 = datetime.datetime(i[0],i[1],i[2])
#     d2 = datetime.datetime(year,i[3],i[4])
#     days = ( d2 - d1 ).days
#     print(days)