package leetcode

import (
	"github.com/panenming/leetcode/datastruct"
)

type TreeNode = datastruct.TreeNode

func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	path := []int{}
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, level, sum int) {
		if node == nil {
			return
		}

		//根据level来更新path
		if level >= len(path) {
			path = append(path, node.Val)
		} else {
			path[level] = node.Val
		}

		sum -= node.Val

		// 到达leaf 并且sum符合条件
		if node.Left == nil && node.Right == nil && sum == 0 {
			temp := make([]int, level+1)
			copy(temp, path)
			res = append(res, temp)
		}
		dfs(node.Left, level+1, sum)
		dfs(node.Right, level+1, sum)
	}

	dfs(root, 0, sum)
	return res
}
