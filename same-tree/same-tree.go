package leetcode

import (
	"github.com/panenming/leetcode/datastruct"
)

type TreeNode = datastruct.TreeNode

func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
