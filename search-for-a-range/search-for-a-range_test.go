package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	one []int
	two int
}

type ans struct {
	one []int
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			para{[]int{5, 7, 7, 8, 8, 10}, 8},
			ans{[]int{3, 4}},
		},

		question{
			para{[]int{5, 7, 7, 8, 8, 10}, 6},
			ans{[]int{-1, -1}},
		},

		question{
			para{[]int{5, 7, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 10}, 8},
			ans{[]int{3, 12}},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, searchRange(p.one, p.two), "输入:%v", p)
	}
}
