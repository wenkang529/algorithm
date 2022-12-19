package solution

import "fmt"

/**
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那 两个 整数，并返回它们的数组下标。 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
 你可以按任意顺序返回答案。 示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
 示例 2：
输入：nums = [3,2,4], target = 6
输出：[1,2]
 示例 3：
输入：nums = [3,3], target = 6
输出：[0,1]
 提示：
 2 <= nums.length <= 10⁴
 -10⁹ <= nums[i] <= 10⁹
 -10⁹ <= target <= 10⁹
 只会存在一个有效答案
 进阶：你可以想出一个时间复杂度小于 O(n²) 的算法吗？
 👍 14368 👎 0

题解：典型的空间换时间，利用hashMap，可以直接O(1)查找。将值存起来，然后查找，最终可以O(n)复杂度解决问题。
*/

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums)-1)
	for index, i := range nums {
		if value, exist := m[i]; exist {
			return []int{value, index}
		}
		m[target-i] = index
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

func RuntwoSum() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}