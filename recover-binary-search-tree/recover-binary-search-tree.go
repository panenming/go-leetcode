package leetcode

import (
	"github.com/panenming/leetcode/datastruct"
)

type TreeNode = datastruct.TreeNode

// 已经假设 BST中没有重复的值并且 root != nil
func recoverTree(root *TreeNode) {
	var first, second, pre *TreeNode
	var dfs func(*TreeNode)

	dfs = func(root *TreeNode) {
		if root.Left != nil {
			dfs(root.Left)
		}

		if pre != nil && pre.Val > root.Val {
			// 如果要调整 [1, 2, 6, 4, 5, 3, 7] 中错位的 6 和 3
			// 其实是把 [6, 4] 中的较大值与 [5, 3] 中的较小值交换。这时，有两组错序。
			// 但是，还有可能是
			// [1, 3, 2] 中错位的 3 和 2，只有一组错序的 [3, 2]
			// 以下的两个 if 语句，可以兼容以上两种情况
			if first == nil {
				first = pre
			}
			if first != nil {
				// 当存在第二组错序的时候
				// second 的值，会被修改
				second = root
			}
		}

		pre = root
		if root.Right != nil {
			dfs(root.Right)
		}
	}
	dfs(root)

	// 题目已经保证存在被交换的节点了
	// 无需检查 first 和 second 是否为 nil
	first.Val, second.Val = second.Val, first.Val
}
