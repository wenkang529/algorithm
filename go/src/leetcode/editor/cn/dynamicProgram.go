package cn

// 动态规划本质上也是分治和循环，就是要找到那个公式（状态转移方程）一般为 f(x+1) = mf(x)... 这种
// 极为下一个状态和上一个状态的关系


/*
一个机器人位于一个 m x n网格的左上角 （起始点在下图中标记为 “Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
问总共有多少条不同的路径？
*/

//https://houbb.github.io/2020/01/23/data-struct-learn-07-base-dp

func uniquePaths(m int, n int) int {
	pathCount := make([][]int, m)
	for i := 0; i < m; i++ {
		pathCount[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				pathCount[i][j] = 1
			} else {
				pathCount[i][j] = pathCount[i-1][j] + pathCount[i][j-1]
			}
		}
	}
	return pathCount[m-1][n-1]
}
