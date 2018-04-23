package leetcode

import (
	"github.com/panenming/leetcode/datastruct"
)

type ListNode = datastruct.ListNode

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	secHead := split(head)
	return merge(sortList(head), sortList(secHead))
}

func split(head *ListNode) *ListNode {
	slow, fast := head, head
	var tail *ListNode
	for fast != nil && fast.Next != nil {
		tail = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 斩断 list
	tail.Next = nil
	return slow
}

func merge(left, right *ListNode) *ListNode {
	var head, cur, pre *ListNode
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur = left
			left = left.Next
		} else {
			cur = right
			right = right.Next
		}

		if head == nil {
			head = cur
		} else {
			pre.Next = cur
		}
		pre = cur
	}

	if left == nil {
		pre.Next = right
	} else {
		pre.Next = left
	}
	return head
}
