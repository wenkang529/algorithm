package cn

/**
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。



 示例 1：


输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]


 示例 2：


输入：head = [1], n = 1
输出：[]


 示例 3：


输入：head = [1,2], n = 1
输出：[1]




 提示：


 链表中结点的数目为 sz
 1 <= sz <= 30
 0 <= Node.val <= 100
 1 <= n <= sz




 进阶：你能尝试使用一趟扫描实现吗？
 👍 2057 👎 0

题解：
对于链表的问题，可以考虑引入一个哑节点，哑节点指向头节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{0, head}
    first, second := head, dummy
    for i := 0; i < n; i++ {
        first = first.Next
    }
    for ; first != nil; first = first.Next {
        second = second.Next
    }
    second.Next = second.Next.Next
    return dummy.Next
}
*/

//leetcode submit region begin(Prohibit modification and deletion)

// Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	root, cur, first := head, head, head
	n-=1
	for n > 0 {
		cur = cur.Next
		n -= 1
	}
	if cur.Next == nil {
		return root.Next
	} else {
		cur = cur.Next
	}
	for cur.Next != nil {
		cur = cur.Next
		first = first.Next
	}
	first.Next = first.Next.Next
	return root
}

//leetcode submit region end(Prohibit modification and deletion)

func RunRemoveMid() {
	root := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: nil,
					},
				},
			},
		},
	}
	removeNthFromEnd(&root, 2)
}