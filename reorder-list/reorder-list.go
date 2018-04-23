package leetcode

import (
	"github.com/panenming/leetcode/datastruct"
)

type ListNode = datastruct.ListNode

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	cur := head
	size := 0

	for cur != nil {
		cur = cur.Next
		size++
	}

	cur = head
	for i := 0; i < (size-1)/2; i++ {
		cur = cur.Next
	}

	next := cur.Next
	for next != nil {
		temp := next.Next
		next.Next = cur
		cur = next
		next = temp
	}

	end := cur
	for head != end {
		hNext := head.Next
		eNext := end.Next
		head.Next = end
		end.Next = hNext
		head = hNext
		end = eNext
	}

	// 封闭 list, 避免出现环状 list
	end.Next = nil
}
