package cn

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序非递归遍历
func preorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var result []int
	for root != nil || len(stack) > 0 {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root.Left)
			root = root.Left
		}
		if len(stack) > 0 {
			root = stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]
		}
	}
	return result
}

// 中序非递归遍历
func inorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var result []int
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) > 0 {
			cur := stack[len(stack)-1]
			result = append(result, cur.Val)
			root = cur.Right
			stack = stack[:len(stack)-1]
		}
	}
	return result
}

// 后序遍历非递归
func postorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var result []int
	pre := &TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == pre {
			result = append(result, root.Val) // 当没有右节点的时候，当前值可以加入答案
			pre = root                        // 避免再次遍历右节点
			root = nil                        // 避免上面的循环遍历左节点
		} else {
			stack = append(stack, root)
			root = root.Right
		}

	}
	return result
}

// 二叉树重建，根据前序遍历 + 中序遍历
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 采用分而治之的思想
	// 前序遍历的第一个值是root，root在中序遍历的位置为index，那么index左边都是左子树，右边就是右子树
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}

// 搜索二叉树的插入
// 插入可以直接找到叶子节点进行插入，不要在中间插入，要在最下面

// 搜索二叉树删除
// 只有当要删除的节点，左右子节点都有的时候才比较复杂
// 查找要删除节点的 右子树的最左节点，然后交换，删除。
// 这时候删的时候变成了只有一个右子节点或者没有子节点，就简单了

// 二叉树剪枝
/*
给定一个二叉树 根节点root，树的每个节点的值要么是 0，要么是 1。请剪除该二叉树中所有节点的值为 0 的子树。
节点 node 的子树为node 本身，以及所有 node的后代。
*/
// 使用递归，这个题目说的是这个节点的所有子树都为0，才应该剪枝
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		root = nil
	}
	return root
}




// 二叉树的层序遍历，运用广度优先搜索的方式
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var l []*TreeNode
	l = append(l, root)
	for len(l) > 0 {
		var tmp []*TreeNode
		var tmpRes []int
		for _, i := range l {
			tmpRes = append(tmpRes, i.Val)
		}
		res = append(res, tmpRes)
		for _, i := range l {
			if i.Left != nil {
				tmp = append(tmp, i.Left)
			}
			if i.Right != nil {
				tmp = append(tmp, i.Right)
			}
		}
		l = tmp
	}
	return res
}
