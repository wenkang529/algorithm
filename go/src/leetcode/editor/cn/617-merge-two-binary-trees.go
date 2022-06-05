package cn
/**
给你两棵二叉树： root1 和 root2 。 

 想象一下，当你将其中一棵覆盖到另一棵之上时，两棵树上的一些节点将会重叠（而另一些不会）。你需要将这两棵树合并成一棵新二叉树。合并的规则是：如果两个节点重叠，那
么将这两个节点的值相加作为合并后节点的新值；否则，不为 null 的节点将直接作为新二叉树的节点。 

 返回合并后的二叉树。 

 注意: 合并过程必须从两个树的根节点开始。 

 

 示例 1： 

 
输入：root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
输出：[3,4,5,5,4,null,7]
 

 示例 2： 

 
输入：root1 = [1], root2 = [1,2]
输出：[2,2]
 

 

 提示： 

 
 两棵树中的节点数目在范围 [0, 2000] 内 
 -10⁴ <= Node.val <= 10⁴ 
 
 👍 989 👎 0

题解：
递归不用都单独再写个子函数，直接用一个函数就行
针对这个题的场景，其实就是谁有就把谁给root1， 两个都有就相加
有个更简单的写法
func mergeTrees(t1, t2 *TreeNode) *TreeNode {
    if t1 == nil {
        return t2
    }
    if t2 == nil {
        return t1
    }
    t1.Val += t2.Val
    t1.Left = mergeTrees(t1.Left, t2.Left)
    t1.Right = mergeTrees(t1.Right, t2.Right)
    return t1
}

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/merge-two-binary-trees/solution/he-bing-er-cha-shu-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func TravelTree(root1 *TreeNode, root2 *TreeNode){
	if root1 == nil && root2 == nil {
		return
	}
	if root1 != nil && root2 != nil {
		root1.Val += root2.Val
		if root1.Left == nil && root2.Left != nil {
			root1.Left = &TreeNode{0, nil, nil}
		}
		if root1.Right == nil && root2.Right != nil {
			root1.Right = &TreeNode{0, nil, nil}
		}
	}
	if root2 != nil {
		TravelTree(root1.Left, root2.Left)
		TravelTree(root1.Right, root2.Right)
	}
}


func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	TravelTree(root1, root2)
	return root1
}
//leetcode submit region end(Prohibit modification and deletion)
