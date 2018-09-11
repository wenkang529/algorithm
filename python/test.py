# import sys

# if __name__ == "__main__":
#     # 读取第一行的n
#     datas=[]
#     lines = int(sys.stdin.readline().strip())
#     for i in range(lines):
#         line = sys.stdin.readline().strip()
#         m,n=map(int, line.split())
#         datas.append([m,n])
#     r1=[]
#     for data in datas:
#         m,n=data[0],data[1]
#         if m>n:
#             m,n=n,m
#         if m>=3:
#             r1.append((m-2)*(n-2))
#             continue
#         if m==1:
#             if n>1:
#                 r1.append(n-2)
#             else:
#                 r1.append(1)
#             continue
#         if m==2:
#             r1.append(0)
#             continue
#         if m==0:
#             r1.append()
#             continue
#     for i in r1:
#         print(i)
# 
# 
#    

import sys

if __name__ == "__main__":
    # 读取第一行的n
    datas=[]
    lines = sys.stdin.readline().strip()
    num,cmd=lines.split()
    num,cmd=int(num),int(cmd)
    nums=[]
    cmds=[]
    for i in range(num):
        lines = int(sys.stdin.readline().strip())
        nums.append(lines)
    for i in range(cmd):
        line = sys.stdin.readline().strip()
        cmds.append(list(map(int, line.split())))
    new_datas=[0]*len(datas)

    print(cmds)
    print(datas)

    for cmd in cmds:
        if cmd[0]==1:
            print(new_datas[cmd[1]])
        if cmd[0]==2:
            new_datas[cmd[2]]+=cmd[3]
            j=cmd[2]
            while j<len(datas)-1:
                if new_datas[j]>datas[j]:
                    aa=new_datas[j]-datas[j]
                    new_datas[j]=datas[j]
                    new_datas[j+1]+=aa
                    j+=1
                    if j==len(datas)-1 and new_datas[j]>datas[j]:
                        new_datas[j]=datas[j]
                else:
                    break
        