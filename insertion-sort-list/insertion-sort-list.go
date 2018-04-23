package leetcode

import (
	"github.com/panenming/leetcode/datastruct"
)

type ListNode = datastruct.ListNode

func insertionSortList(head *ListNode) *ListNode {
	headPre := &ListNode{Next: head}

	cur := head
	for cur != nil && cur.Next != nil {
		p := cur.Next
		if cur.Val <= p.Val {
			cur = p
			continue
		}

		cur.Next = p.Next
		pre, next := headPre, headPre.Next
		for next.Val < p.Val {
			pre = next
			next = next.Next
		}
		pre.Next = p
		p.Next = next
	}

	return headPre.Next
}
