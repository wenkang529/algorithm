package solution

import "fmt"

/*
1.
给定两个整数 a 和 b ，求它们的除法的商 a/b ，要求不得使用乘号 '*'、除号 '/' 以及求余符号 '%'。

注意：

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231, 231−1]。本题中，如果除法结果溢出，则返回 231 − 1
*/

// 计算a/b 返回b左移多少位可以刚好不大于a
// 由于为了方便计算，都采用了负数（最大的负数直接转正数会越界）

// 计算a/b 返回b左移多少位可以刚好不大于a
// 由于为了方便计算，都采用了负数（最大的负数直接转正数会越界）
func quickDiv(a, b int) int {
	count := 0
	for i := 0; i < 32; i++ {
		if a-b<<i > 0 {
			break
		}
		count = i
	}
	return count
}

func divide(a int, b int) int {
	if b == 0 || a == 0 {
		return 0
	}
	if b == -1 && a == (0-1<<31) {
		return 1<<31 - 1
	}
	flag := false
	if a > 0 && b < 0 {
		a = 0 - a
		flag = true
	}
	if a < 0 && b > 0 {
		b = 0 - b
		flag = true
	}
	if a > 0 && b > 0 {
		a = 0 - a
		b = 0 - b
	}

	res := 0
	for a <= b {
		x := quickDiv(a, b)
		a = a - b<<x
		res += 1 << x
	}
	if flag {
		res = 0 - res
	}
	return res
}

/*给定两个 01 字符串 a 和 b ，请计算它们的和，并以二进制字符串的形式输出。

输入为 非空 字符串且只包含数字 1 和 0。

 

示例 1:

输入: a = "11", b = "10"
输出: "101"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"
 

提示：

每个字符串仅由字符 '0' 或 '1' 组成。
1 <= a.length, b.length <= 10^4
字符串如果不是 "0" ，就都不含前导零。
 
*/

func addBinary(a string, b string) string {
	lena := len(a)
	lenb := len(b)
	if lena < lenb {
		a, b = b, a
		lena, lenb = lenb, lena
	}
	res := make([]byte, lena+1)
	for i := 0; i < lena; i++ {
		m := lena - i - 1
		n := lenb - i - 1
		bi := byte('0')
		if n >= 0 {
			bi = b[n]
		}
		ai := a[m]
		if ai == bi && ai == '1' {
			res[i+1] = '1'
		}
		if ai != bi {
			if res[i] == '1' {
				res[i+1] = '1'
				res[i] = '0'
			} else {
				res[i] = '1'
			}
		}
	}
	result := ""
	if res[lena] == '1' {
		result = "1"
	}
	for i := lena - 1; i >= 0; i-- {
		if res[i] == '1' {
			result += "1"
		} else {
			result += "0"
		}
	}
	return result
}

func maxProduct(words []string) int {
	res := 0
	indexMap := map[int64]int{}
	for _, word := range words {
		key := int64(0)
		for _, w := range word {
			key |= 1 << (w - 'a' + 1)
		}
		indexMap[key] = len(word)
	}
	for key1, value1 := range indexMap {
		for key2, value2 := range indexMap {
			if key1&key2 == 0 && value1*value2 > res {
				res = value1 * value2
			}
		}
	}
	return res
}

func Rungogogo() {
	fmt.Println(maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}))
	//fmt.Println()
}
