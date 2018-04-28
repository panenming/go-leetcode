package leetcode

import (
	"github.com/panenming/leetcode/datastruct"
)

type ListNode = datastruct.ListNode

func removeElements(root *ListNode, val int) *ListNode {
	headPre := &ListNode{Next: root}
	temp := headPre
	for temp.Next != nil {
		if temp.Next.Val == val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}

	return headPre.Next
}
