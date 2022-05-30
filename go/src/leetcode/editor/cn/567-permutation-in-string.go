package cn

import "fmt"

/**
给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。

 换句话说，s1 的排列之一是 s2 的 子串 。



 示例 1：


输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").


 示例 2：


输入：s1= "ab" s2 = "eidboaoo"
输出：false




 提示：


 1 <= s1.length, s2.length <= 10⁴
 s1 和 s2 仅包含小写字母

 👍 685 👎 0

题解：
误区1： 进入了思维死胡同，感觉越来越复杂
误区2： 想要初始化词典的时候写的 mDict := map[byte]int  不能是 ：=
误区2： 其实不需要复制词典，只需要left向前回退就行了


func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	mDict := map[byte]int{}
	mDictBack := map[byte]int{}
	for index := range s1 {
		if _, ok := mDict[s1[index]]; ok {
			mDict[s1[index]] ++
		} else {
			mDict[s1[index]] = 1
		}
	}
	for k, v := range mDict {
		mDictBack[k] = v
	}
	first := 0
	for index := range s2 {
		value := s2[index]
		if v, ok := mDict[value]; ok {
			if v > 0 {
				mDict[value] --
				if index - first == len(s1) -1 {
					return true
				}
			} else {
				for first <= index {
					if s2[first] == value {
						first ++
						break
					}
					mDict[s2[first]] ++
					first ++
				}
			}
		} else {
			first = index + 1
			for k, v := range mDictBack {
				mDict[k] = v
			}
		}
	}
	return false
}

*/

//leetcode submit region begin(Prohibit modification and deletion)
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	mDict := map[byte]int{}
	for index := range s1 {
		if _, ok := mDict[s1[index]]; ok {
			mDict[s1[index]] ++
		} else {
			mDict[s1[index]] = 1
		}
	}
	first := 0
	for index := range s2 {
		value := s2[index]
		if v, ok := mDict[value]; ok && v > 0 {
			mDict[value] --
			if index-first == len(s1)-1 {
				return true
			}
		} else {
			for first <= index {
				if s2[first] == value {
					first++
					break
				}
				mDict[s2[first]] ++
				first++
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func RunInString(s1, s2 string) {
	fmt.Println(checkInclusion(s1, s2))

}
