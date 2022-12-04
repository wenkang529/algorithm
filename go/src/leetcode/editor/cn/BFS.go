package cn

// 广度优先搜索
/*
给你一个由'1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。
 */

// 广度优先搜索，把走过的都进行染色，避免重复走

func numIslands(grid [][]byte) int {
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
				var subList [][]int
				subList = append(subList, []int{i, j})
				for len(subList) > 0 {
					for _, i := range subList {
						x, y := i[0], i[1]
						grid[x][y] = '0'
						subList = subList[1:]
						if x+1 < m && grid[x+1][y] == '1' {
							grid[x+1][y] = '0'
							subList = append(subList, []int{x + 1, y})
						}
						if x-1 >= 0 && grid[x-1][y] == '1' {
							grid[x-1][y] = '0'
							subList = append(subList, []int{x - 1, y})
						}
						if y+1 < n && grid[x][y+1] == '1' {
							grid[x][y+1] = '0'
							subList = append(subList, []int{x, y + 1})
						}
						if y-1 >= 0 && grid[x][y-1] == '1' {
							grid[x][y-1] = '0'
							subList = append(subList, []int{x, y - 1})
						}
					}
				}
			}
		}
	}
	return islandCount
}