package leetcode

import (
	"fmt"
	"testing"

	"github.com/panenming/leetcode/datastruct"
	"github.com/stretchr/testify/assert"
)

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	// tcs is testcase slice
	tcs := []struct {
		pre, in []int
		ans     int
	}{

		{
			[]int{},
			[]int{},
			0,
		},

		{
			[]int{-10, 100, -1, -2},
			[]int{-10, -1, 100, -2},
			100,
		},

		{
			[]int{10, 10, -1, -2},
			[]int{10, -1, 10, -2},
			20,
		},

		{
			[]int{1, 1, 1, 1, 1, 10, 10},
			[]int{1, 1, 1, 1, 10, 1, 10},
			21,
		},

		{
			[]int{1, 2, 3},
			[]int{2, 1, 3},
			6,
		},

		{
			[]int{4, 2, 1, 3, 6, 5, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
			22,
		},

		// 可以多个 testcase
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		root := datastruct.PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.ans, maxPathSum(root), "输入:%v", tc)
	}
}
