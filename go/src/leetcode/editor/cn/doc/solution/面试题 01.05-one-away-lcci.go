package solution

import (
	"fmt"
	"math"
)

/**
字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。



 示例 1:

 输入:
first = "pale"
second = "ple"
输出: True



 示例 2:

 输入:
first = "pales"
second = "pal"
输出: False

 👍 176 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func oneEditAway(first string, second string) bool {
	if math.Abs(float64(len(first)-len(second))) > 1 {
		return false
	}
	if len(first) == len(second) {
		if first == second {
			return true
		}
		flag := 0
		for i := 0; i < len(first); i++ {
			if first[i] != second[i] {
				if flag == 1 {
					return false
				}
				flag += 1
			}
		}
		return true
	}
	if len(first) > len(second) {
		first, second = second, first
	}
	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			if first[i:] != second[i+1:] {
				return false
			}
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func RunoneEditAway() {
	fmt.Println(oneEditAway("palea", "pales"))

}
