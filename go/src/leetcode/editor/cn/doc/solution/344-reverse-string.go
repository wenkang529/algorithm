package solution
/**
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。 

 不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。 

 

 示例 1： 

 
输入：s = ["h","e","l","l","o"]
输出：["o","l","l","e","h"]
 

 示例 2： 

 
输入：s = ["H","a","n","n","a","h"]
输出：["h","a","n","n","a","H"] 

 

 提示： 

 
 1 <= s.length <= 10⁵ 
 s[i] 都是 ASCII 码表中的可打印字符 
 
 👍 595 👎 0

题解：
比较简单，但是有两个优化的地方
1。 golang交换值的时候不用定一个tmp，可以直接 a,b = b, a
2. 可以使用left，right指针，不用判断len/2

*/

//leetcode submit region begin(Prohibit modification and deletion)
func reverseString(s []byte)  {
	for index, v := range s {
		if index >= len(s)/2 {
			break
		}
		tmp := v
		s[index] = s[len(s)-index-1]
		s[len(s)-index-1]=tmp
	}

}
//leetcode submit region end(Prohibit modification and deletion)
