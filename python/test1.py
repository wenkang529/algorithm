# #!/bin/python
# # -*- coding: utf8 -*-
# import sys
# import os
# import re


# # 请完成下面这个函数，实现题目要求的功能
# # 当然，你也可以不按照下面这个模板来作答，完全按照自己的想法来 ^-^ 
# # ******************************开始写代码******************************


# def GetResult(K):
#     r=[]
#     for i in range(10000,30001):
#         a=i//100
#         if(a%K!=0):
#             continue
#         b=(i//10-(i//10000)*1000)
#         c=i-(a//10)*1000


#         if(b%K==0 and c%K==0):
#             r.append(i)
#     return r



# _K = int(input())
# res=GetResult(_K)
# if len(res)==0:
#     print('NO')
# else:
#     for i in res:
#         print(i)


while(True):
    arr=list(map(int,list(input().split())))
    if(len(arr)==0):
        break
    print(arr)
    a,b,c,n=arr
    l=[]
    for i in range(n):
        l.append(input())
    print(l)
    for ll in l:
        ll=list(map(int,list(ll)))
        x1,y1,x2,y2,x3,y3,x4,y4=ll
        print(x1,y1,x2,y2,x3,y3,x4,y4)
        if(x1==x2):
            r1=x1
            r2=x3
        else:
            r1=x1
            r2=x2
        ry1=(-1*a/(b*1.0)*r1-c)
        ry2=(-1*a/(b*1.0)*r2-c)
        if(ry1>=y1):
            if(ry1>=y2 and ry1>=y3 and ry1>=y4):
                print(1)
            else:
                print(-1)
        else:
            if(ry1<=y2 and ry1<=y3 and ry1<=y4):
                print(1)
            else:
                print(-1)
                

