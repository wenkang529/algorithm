package cn

import "fmt"

/**
有一幅以 m x n 的二维整数数组表示的图画 image ，其中 image[i][j] 表示该图画的像素值大小。

 你也被给予三个整数 sr , sc 和 newColor 。你应该从像素 image[sr][sc] 开始对图像进行 上色填充 。

 为了完成 上色工作 ，从初始像素开始，记录初始坐标的 上下左右四个方向上 像素值与初始坐标相同的相连像素点，接着再记录这四个方向上符合条件的像素点与他们对应
四个方向上 像素值与初始坐标相同的相连像素点，……，重复该过程。将所有有记录的像素点的颜色值改为 newColor 。

 最后返回 经过上色渲染后的图像 。



 示例 1:




输入: image = [[1,1,1],[1,1,0],[1,0,1]]，sr = 1, sc = 1, newColor = 2
输出: [[2,2,2],[2,2,0],[2,0,1]]
解析: 在图像的正中间，(坐标(sr,sc)=(1,1)),在路径上所有符合条件的像素点的颜色都被更改成2。
注意，右下角的像素没有更改为2，因为它不是在上下左右四个方向上与初始点相连的像素点。


 示例 2:


输入: image = [[0,0,0],[0,0,0]], sr = 0, sc = 0, newColor = 2
输出: [[2,2,2],[2,2,2]]




 提示:


 m == image.length
 n == image[i].length
 1 <= m, n <= 50
 0 <= image[i][j], newColor < 2¹⁶
 0 <= sr < m
 0 <= sc < n

 👍 337 👎 0

题解：解法分为深度优先和广度优先

广度优先算法：就是把上下左右分别都入栈，然后遍历堆栈并处理，并把新的上下左右入栈。如下
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	lenr, lenc := len(image), len(image[0])
	originColor := image[sr][sc]
	if image[sr][sc] == newColor {
		return image
	}
	var queue [][]int
	queue = append(queue, []int{sr, sc})
	for i := 0; i < len(queue); i++ {
		sr, sc = queue[i][0], queue[i][1]
		image[sr][sc] = newColor
		if sr-1 >= 0 && image[sr-1][sc] == originColor {
			queue = append(queue, []int{sr - 1, sc})
		}
		if sr+1 < lenr && image[sr+1][sc] == originColor {
			queue = append(queue, []int{sr + 1, sc})
		}
		if sc-1 >= 0 && image[sr][sc-1] == originColor {
			queue = append(queue, []int{sr, sc - 1})
		}
		if sc+1 < lenc && image[sr][sc+1] == originColor {
			queue = append(queue, []int{sr, sc + 1})
		}
	}
	return image
}

深度优先，就是典型的二叉树的前/中/后序遍历
其实深度遍历更简单，就是递归，递归的一个重要特性是，新的加入的可以和原来的同样的判断处理逻辑

func deal(image [][]int, sr, sc int, origin, newColor int) {
	if image[sr][sc] != origin {
		return
	}
	image[sr][sc] = newColor
	if sr-1 >= 0 {
		deal(image, sr-1, sc, origin, newColor)
	}
	if sr+1 < len(image) {
		deal(image, sr+1, sc, origin, newColor)
	}
	if sc-1 >= 0 {
		deal(image, sr, sc-1, origin, newColor)
	}
	if sc+1 < len(image[0]) {
		deal(image, sr, sc+1, origin, newColor)
	}
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}
	deal(image, sr, sc, image[sr][sc], newColor)
	return image
}

*/
func deal(image [][]int, sr, sc int, origin, newColor int) {
	if image[sr][sc] != origin {
		return
	}
	image[sr][sc] = newColor
	if sr-1 >= 0 {
		deal(image, sr-1, sc, origin, newColor)
	}
	if sr+1 < len(image) {
		deal(image, sr+1, sc, origin, newColor)
	}
	if sc-1 >= 0 {
		deal(image, sr, sc-1, origin, newColor)
	}
	if sc+1 < len(image[0]) {
		deal(image, sr, sc+1, origin, newColor)
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}
	deal(image, sr, sc, image[sr][sc], newColor)
	return image
}

//leetcode submit region end(Prohibit modification and deletion)

func RunFloodFill() {
	res := floodFill([][]int{
		{1, 1, 1}, {1, 1, 0}, {1, 0, 1},
	}, 1, 1, 2)
	fmt.Println(res)
}
