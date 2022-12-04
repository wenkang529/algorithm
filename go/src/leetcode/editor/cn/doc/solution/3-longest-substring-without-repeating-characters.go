package solution

import (
	"fmt"
)

/**
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。



 示例 1:


输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。


 示例 2:


输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。


 示例 3:


输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。




 提示：


 0 <= s.length <= 5 * 10⁴
 s 由英文字母、数字、符号和空格组成

 👍 7641 👎 0


// 题解： 我的第一次的方法是采用了两次循环的方法，其实其次循环就可以，采用滑动窗口的方式
func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	max := 0
	tDict := map[uint8]byte{}
	for i, _ := range s {
		for index := i; index < len(s); index++ {
			if _, ok := tDict[s[index]]; ok {
				if len(tDict) > max {
					max = len(tDict)
				}
				tDict = map[uint8]byte{}
				break
			} else {
				tDict[s[index]] = 0
			}
		}
	}
	return max
}

滑动窗口其实就是左右指针如何移动的问题

计算时间复杂度的时候，我看代码有两次for循环，以为就是O(N2), 其实不是，是O(n)。 因为删除元素的时候最多会删除n次，所以是2n，即为O(n)

对于字符串，发现如果是 for index, value := range s ， 那么value的类型是int32， 但是 s[index] 就是byte 。。。

*/

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	max := 0
	first := 0
	tDict := map[byte]int{}
	for index := range s {
		if value, ok := tDict[s[index]]; ok {
			if index-first > max {
				max = index - first
			}
			for first <= value {
				delete(tDict, s[first])
				first++
			}
		}
		tDict[s[index]] = index
	}
	if len(tDict) > max {
		max = len(tDict)
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)

func RunlongestString() {
	fmt.Println(lengthOfLongestSubstring("jbpnbwwd"))
}
