package leetcode

import (
	"github.com/panenming/leetcode/datastruct"
)

type TreeNode = datastruct.TreeNode

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	l := rightSideView(root.Left)
	r := rightSideView(root.Right)

	if len(l) > len(r) {
		r = append(r, l[len(r):]...)
	}

	res := make([]int, len(r)+1)
	res[0] = root.Val
	copy(res[1:], r)
	return res
}
