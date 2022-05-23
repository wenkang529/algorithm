package cn

/**
由范围 [0,n] 内所有整数组成的 n + 1 个整数的排列序列可以表示为长度为 n 的字符串 s ，其中:
 如果 perm[i] < perm[i + 1] ，那么 s[i] == 'I'
 如果 perm[i] > perm[i + 1] ，那么 s[i] == 'D'
 给定一个字符串 s ，重构排列 perm 并返回它。如果有多个有效排列perm，则返回其中 任何一个 。
 示例 1：
输入：s = "IDID"
输出：[0,4,1,3,2]
 示例 2：
输入：s = "III"
输出：[0,1,2,3]
 示例 3：
输入：s = "DDI"
输出：[3,2,0,1]
 提示：
 1 <= s.length <= 10⁵
 s 只包含字符 "I" 或 "D"
*/

/*
解题思路：
贪心算法。不关心全局最优解，只关注局部最优解，
先把眼前最外层问题解决了，里面一层的问题类似于最外层的问题，再逐步解决。

Tips：
- 先解决第一个，下一个问题是不是就变成了和上一个同样的问题呢？是的话就可以用贪心算法。
- 注意边界条件，在最后一个不要遗漏了
- 既然是固定长度，使用make的时候可以直接申请固定的长度。


*/

//leetcode submit region begin(Prohibit modification and deletion)
func diStringMatch(s string) []int {
	n := len(s)
	res := make([]int, n+1)
	low, high := 0, n
	for index, i := range s {
		if i == 'I' {
			res[index] = low
			low += 1
		}
		if i == 'D' {
			res[index] = high
			high -= 1
		}
	}
	res[n] = low
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func RundiStringMatch() {
	diStringMatch("IDIDD")
}
