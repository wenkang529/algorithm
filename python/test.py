#!/bin/python
# -*- coding: utf8 -*-
import numpy as np

m = int(input())
n = int(input())
arr=list(input().split())
arr=map(int,arr)
arr=list(arr)
arr=np.array(arr)
arr=np.reshape(arr,(m,n))
a=b=0
c=m-1
d=n-1

while(a!=c or b!=d):
    for i in range(b,d+1):
        print(str(arr[a][i]),'\n')
    for i in range(a+1,c+1):
        print(str(arr[i][d]),'\n')
    for i in range(b,d):
        print(str(arr[c][d-1-i]),'\n')
    for i in range(c-1,a,-1):
        print(str(arr[i][b]),'\n')
    a+=1
    b+=1
    c-=1
    d-=1
if(a==c and b==d):
    print(arr[a][b])