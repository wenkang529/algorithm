def numIslands(grid):
    count = 0
    r1=0
    r2=0
    def sink(i, j,count):
        if 0 <= i < len(grid) and 0 <= j < len(grid[i]) and grid[i][j] == 1:
            grid[i][j] = '0'
            count+=1
            sink(i,j,count)
            sink(i+1, j, count)
            sink(i-1, j, count)
            sink(i, j+1, count)
            sink(i, j-1, count)
            sink(i+1, j+1, count)
            sink(i-1,j-1,count)
            return 1,None
        return 0,count
    for i in range(len(grid)):
        for j in range(len(grid[i])):
            flag,num=sink(i,j,count)
            # print(flag,num)
            if num is not None:
                r1+=1
                if num>r2:
                    r2=num
    return r1,r2











# import sys
# if __name__ == "__main__":
#     # 读取第一行的n
#     n = int(sys.stdin.readline().strip())
#     for i in range(n):
#         # 读取每一行
#         line = sys.stdin.readline().strip()
#         # 把每一行的数字分隔后转化成int列表
#         values = list(map(int, line.split()))
#     r1,r2=numIslands(values)
#     print(r1,",",r2)
a=[[1,0,1],[0,0,0]]
print(numIslands(a))

