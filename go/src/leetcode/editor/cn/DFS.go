package cn

// 深度优先搜索，用到递归函数（或者可以用循环？）

/*
给你一个由'1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。
*/

func NumIslandsWithDFS(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	islandCount := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				islandCount += 1
				inArea(grid, i, j)
			}
		}
	}
	return islandCount
}

func inArea(grid [][]byte, i, j int) {
	if i >= len(grid) || i < 0 || j >= len(grid[0]) || j < 0 {
		return
	}
	if grid[i][j] != '1' {
		return
	}
	grid[i][j] = '0'
	inArea(grid, i+1, j)
	inArea(grid, i-1, j)
	inArea(grid, i, j+1)
	inArea(grid, i, j-1)
}
