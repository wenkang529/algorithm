#coding=utf-8
import sys
if __name__ == "__main__":
    # 读取第一行的n
    n = int(sys.stdin.readline().strip())
    arr=[]
    for i in range(n):
        # 读取每一行
        line = sys.stdin.readline().strip()
        # 把每一行的数字分隔后转化成int列表
        values = list(map(lambda x:int(x)-1, line.split()))
        values.pop(-1)
        arr.append(values)
    result=0
    is_find=[0]*n
    sets=[]
    print(arr)


    while(sum(is_find)<n):
        new=set()
        first_flag=1
        for index,i in enumerate(arr):
            if is_find[index]==1:
                continue
            else:
                is_find[index]=1
            if first_flag==1:
                new.update(i)
                is_find[index]=1
                continue
            for j in i:
                if j in new:
                    new.update(i)
                    break
        result+=1
    print(result)

