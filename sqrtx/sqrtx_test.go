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
	x int
}

// ans 是答案
type ans struct {
	one int
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				15,
			},
			ans{
				3,
			},
		},

		question{
			para{
				3,
			},
			ans{
				1,
			},
		},

		question{
			para{
				9,
			},
			ans{
				3,
			},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, mySqrt(p.x), "输入:%v", p)
	}
}
