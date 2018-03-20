# import time
#
# def fib(n):
#     if n==0 or n==1:
#         return n
#     return fib(n-1)+fib(n-2)
#
#
#
# dict = {}
# dict[0]=0
# dict[1]=1
#
# def fib1(n):
#     if n not in dict:
#         dict[n]=fib1(n-1)+fib(n-2)
#     return dict[n]
#
# x=25
#
# s1=time.time()
# print(fib(x))
# e1=time.time()
# print('l1:',e1-s1)
#
#
# s2=time.time()
# print(fib1(x))
# e2=time.time()
# print('l2:',e2-s2)

s='aabbcc'
p='a.b*cc'

table = [[False] * (len(s) + 1) for _ in range(len(p) + 1)]

print(table)