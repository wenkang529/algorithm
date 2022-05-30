package cn

import "fmt"

/**
给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。



 示例 1:


输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]


 示例 2:


输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]



 提示：


 1 <= nums.length <= 10⁵
 -2³¹ <= nums[i] <= 2³¹ - 1
 0 <= k <= 10⁵




 进阶：


 尽可能想出更多的解决方案，至少有 三种 不同的方法可以解决这个问题。
 你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？


 // 解题思路，在函数内通过slice修改指针的起始点是不生效的，因为底层数据也没变，而原来的nums指针也没有变化
这道题的核心是怎么能占用空间少
对比环状替换，采用数组反转的方式会简单很多。


 👍 1486 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func reverse(nums []int) {
	left, right := 0, len(nums)-1
	for left < right {
		t := nums[left]
		nums[left] = nums[right]
		nums[right] = t
		left += 1
		right -= 1
	}
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

//leetcode submit region end(Prohibit modification and deletion)

func RunRotate() {
	nums := []int{-1, -100, 3, 99}
	rotate(nums, 2)
	fmt.Println(nums)
}
