package solution

import (
	"main/src/leetcode/editor/cn"
	"math"
	"strconv"
	"strings"
)

/**
序列化是将数据结构或对象转换为一系列位的过程，以便它可以存储在文件或内存缓冲区中，或通过网络连接链路传输，以便稍后在同一个或另一个计算机环境中重建。

 设计一个算法来序列化和反序列化 二叉搜索树 。 对序列化/反序列化算法的工作方式没有限制。 您只需确保二叉搜索树可以序列化为字符串，并且可以将该字符串反序列化
为最初的二叉搜索树。

 编码的字符串应尽可能紧凑。



 示例 1：


输入：root = [2,1,3]
输出：[2,1,3]


 示例 2：


输入：root = []
输出：[]




 提示：


 树中节点数范围是 [0, 10⁴]
 0 <= Node.val <= 10⁴
 题目数据 保证 输入的树是一棵二叉搜索树。

 👍 333 👎 0
算法背景：

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
* Definition for a binary tree node.

 */

type Codec struct{}

func Constructor() (_ Codec) { return }

func (Codec) serialize(root *cn.TreeNode) string {
	var arr []string
	var postOrder func(*cn.TreeNode)
	postOrder = func(node *cn.TreeNode) {
		if node == nil {
			return
		}
		postOrder(node.Left)
		postOrder(node.Right)
		arr = append(arr, strconv.Itoa(node.Val))
	}
	postOrder(root)
	return strings.Join(arr, " ")
}

func (Codec) deserialize(data string) *cn.TreeNode {
	if data == "" {
		return nil
	}
	arr := strings.Split(data, " ")
	var construct func(int, int) *cn.TreeNode
	construct = func(lower, upper int) *cn.TreeNode {
		if len(arr) == 0 {
			return nil
		}
		val, _ := strconv.Atoi(arr[len(arr)-1])
		if val < lower || val > upper {
			return nil
		}
		arr = arr[:len(arr)-1]
		return &cn.TreeNode{Val: val, Right: construct(val, upper), Left: construct(lower, val)}
	}
	return construct(math.MinInt32, math.MaxInt32)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
//leetcode submit region end(Prohibit modification and deletion)
func Runserialize() {
	root := cn.TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	ser := Constructor()
	deser := Constructor()
	tree := ser.serialize(&root)
	deser.deserialize(tree)
}
