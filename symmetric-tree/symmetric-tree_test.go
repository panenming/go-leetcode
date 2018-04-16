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
		ans     bool
	}{

		{
			[]int{},
			[]int{},
			true,
		},

		{
			[]int{1, 2, 2},
			[]int{2, 1, 2},
			true,
		},

		{
			[]int{1, 2, 3, 2, 3},
			[]int{2, 3, 1, 2, 3},
			false,
		},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, isSymmetric(datastruct.PreIn2Tree(tc.pre, tc.in)), "输入:%v", tc)
	}
}
