package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	d, headIsNthFromEnd := getParent(head, n)
	if headIsNthFromEnd {
		// 删除head节点
		return head.Next
	}

	d.Next = d.Next.Next

	return head
}

func getParent(head *ListNode, n int) (parent *ListNode, headIsNthFromEnd bool) {
	parent = head
	for head != nil {
		// n < 0 表示是从前往后走，当head到达尾部，parent正好是从后往前数n个的parent，经典
		if n < 0 {
			parent = parent.Next
		}

		n--
		head = head.Next

	}
	// n==0，说明链的长度等于n
	headIsNthFromEnd = n == 0
	return
}
