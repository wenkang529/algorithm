package solution

/**
给你一个大小为 m x n 的二进制矩阵 grid 。

 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。你可以假设 grid 的四个边缘都被
0（代表水）包围着。

 岛屿的面积是岛上值为 1 的单元格的数目。

 计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。



 示例 1：


输入：grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,
0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,
0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
输出：6
解释：答案不应该是 11 ，因为岛屿只能包含水平或垂直这四个方向上的 1 。


 示例 2：


输入：grid = [[0,0,0,0,0,0,0,0]]
输出：0




 提示：


 m == grid.length
 n == grid[i].length
 1 <= m, n <= 50
 grid[i][j] 为 0 或 1

 👍 788 👎 0

题解：
典型的深度优先，不要忽略 >=0 ，而不是 > 0

可以采用一个函数内的全局变量作为计数

*/

//leetcode submit region begin(Prohibit modification and deletion)
func numPerIsland(grid [][]int, r, c, count int) int {
	if grid[r][c] == 0 {
		return count
	}
	grid[r][c] = 0
	count += 1
	for _, value := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		newC := r + value[0]
		newR := c + value[1]
		if newC >= 0 && newC < len(grid) && newR >= 0 && newR < len(grid[0]) {
			count = numPerIsland(grid, newC, newR, count)
		}
	}
	return 1
}

func maxAreaOfIsland(grid [][]int) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			c := numPerIsland(grid, i, j, 0)
			if c > count {
				count = c
			}
		}
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)

func RunMaxAreaOfIsland() {
	maxAreaOfIsland([][]int{
		{1, 1},
	})
}
