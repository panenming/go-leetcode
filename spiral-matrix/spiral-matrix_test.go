package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

// para 是参数
type para struct {
	matrix [][]int
}

// ans 是答案
type ans struct {
	one []int
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[][]int{
					[]int{},
					[]int{},
					[]int{},
				},
			},
			ans{
				[]int{},
			},
		},

		question{
			para{
				[][]int{
					[]int{7},
					[]int{9},
					[]int{6},
				},
			},
			ans{
				[]int{7, 9, 6},
			},
		},

		question{
			para{
				[][]int{
					[]int{1, 2, 3},
					[]int{4, 5, 6},
					[]int{7, 8, 9},
				},
			},
			ans{
				[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
			},
		},

		question{
			para{
				[][]int{
					[]int{1, 2, 3},
					[]int{4, 5, 6},
				},
			},
			ans{
				[]int{1, 2, 3, 6, 5, 4},
			},
		},

		question{
			para{
				[][]int{
					[]int{1, 2, 3},
				},
			},
			ans{
				[]int{1, 2, 3},
			},
		},

		question{
			para{
				[][]int{},
			},
			ans{
				[]int{},
			},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, spiralOrder(p.matrix), "输入:%v", p)
	}
}
