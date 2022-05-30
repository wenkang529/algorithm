package cn

import (
	"fmt"
	"strings"
)

/**
给定一个字符串 s ，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。



 示例 1：


输入：s = "Let's take LeetCode contest"
输出："s'teL ekat edoCteeL tsetnoc"


 示例 2:


输入： s = "God Ding"
输出："doG gniD"




 提示：


 1 <= s.length <= 5 * 10⁴
 s 包含可打印的 ASCII 字符。
 s 不包含任何开头或结尾空格。
 s 里 至少 有一个词。
 s 中的所有单词都用一个空格隔开。

 👍 446 👎 0

题解：
string的遍历是byte，可以定义一个[]byte 转为string的时候可以直接 string([]byte{})

如果是[]string 转为一个字符串，可以使用 strings.Join(s, "")

*/

//leetcode submit region begin(Prohibit modification and deletion)
func reverseWords(s string) string {
	left, right := 0, 0
	var res []string
	for index, v := range s {
		if v == ' ' {
			if left != right {
				// reverse
				for right = right-1; right >= left; right-- {
					res = append(res, string(s[right]))
				}
			}
			res = append(res, string(v))
			left, right = index+1, index+1
		} else {
			right = index +1
		}
	}
	if left != right {
		for right=right-1; right>= left; right-- {
			res = append(res, string(s[right]))
		}
	}
	return strings.Join(res, "")
}

//leetcode submit region end(Prohibit modification and deletion)

func RunReverseWords() {
	fmt.Println(reverseWords("I love u"))
}
