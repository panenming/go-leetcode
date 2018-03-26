package leetcode

/**
定义数组node
**/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	temp := result
	v, n := 0, 0
	for {
		v, n = add(l1, l2, n)
		temp.Val = v
		// 下一位
		l1 = next(l1)
		l2 = next(l2)
		// 下一位为nil结束
		if l1 == nil && l2 == nil {
			break
		}

		temp.Next = &ListNode{}
		temp = temp.Next
	}

	if n == 1 {
		temp.Next = &ListNode{Val: n}
	}
	return result
}

func add(n1, n2 *ListNode, i int) (v, n int) {
	if n1 != nil {
		v += n1.Val
	}
	if n2 != nil {
		v += n2.Val
	}
	v += i
	if v > 9 {
		v -= 10
		n = 1
	}
	return
}

func next(l *ListNode) *ListNode {
	if l != nil {
		return l.Next
	}
	return nil
}
