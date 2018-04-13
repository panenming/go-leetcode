package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 存放< x的节点
	lessHead := &ListNode{}
	// 存放 > x的节点
	noLessHead := &ListNode{}

	lessEnd := lessHead
	noLesssEnd := noLessHead

	for head != nil {
		if head.Val < x {
			lessEnd.Next = head
			lessEnd = lessEnd.Next
		} else {
			noLesssEnd.Next = head
			noLesssEnd = noLesssEnd.Next
		}
		head = head.Next
	}

	lessEnd.Next = noLessHead.Next
	noLesssEnd.Next = nil
	head = lessHead.Next
	return head
}
