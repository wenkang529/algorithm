#coding=utf-8
import sys
if __name__ == "__main__":
    # 读取第一行的n
    n = int(sys.stdin.readline().strip())
    data = sys.stdin.readline().strip()

    ac=list(ac.split(';'))
    dic={}
    for i in ac:
        action,value=i.split('_')
        value=list(value.split('|'))
        for j in value:
            if j not in dic.keys():
                tmp=[]
                tmp.append(action)
                dic[j]=tmp
            else:
                dic[j].append(action)

    keys=list(dic.keys())
    keys.sort(key=len)
    for i in keys:
        tmp=''
        for i in dic[i]:
            tmp=tmp+i+','
        tmp=tmp[0:-2]
        dic[i]=tmp
    min_len=len(keys[0])
    max_len=len(keys[-1])
    result=''
    i=0
    while i <len(data):
        flag=1
        for j in range(max_len,min_len-1,-1):
            if i+j>=len:
                result+=data[i:]
                break    
            if data[i,i+j] in dic.keys():
                result=result+data[i,i+j]+'/'+dic[i]
                i=i+j
                flag=0
                break
        if flag:
            result+=data[i:i+j]
            
    print(result)
                


    # arr=[]
    # for i in range(n):
    #     # 读取每一行
    #     line = sys.stdin.readline().strip()
    #     # 把每一行的数字分隔后转化成int列表
    #     values = list(map(lambda x:int(x)-1, line.split()))
    #     values.pop(-1)
    #     arr.append(values)
    # result=0
    # is_find=[0]*n
    # sets=[]
    # print(arr)


    # while(sum(is_find)<n):
    #     new=set()
    #     first_flag=1
    #     for index,i in enumerate(arr):
    #         if is_find[index]==1:
    #             continue
    #         else:
    #             is_find[index]=1
    #         if first_flag==1:
    #             new.update(i)
    #             is_find[index]=1
    #             continue
    #         for j in i:
    #             if j in new:
    #                 new.update(i)
    #                 break
    #     result+=1
    # print(result)

