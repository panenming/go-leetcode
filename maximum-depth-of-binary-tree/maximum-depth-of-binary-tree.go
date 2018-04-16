package leetcode

import (
	"github.com/panenming/leetcode/datastruct"
)

type TreeNode = datastruct.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
