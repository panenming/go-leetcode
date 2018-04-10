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

type para struct {
	one []Interval
}

type ans struct {
	one []Interval
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[]Interval{
				Interval{2, 3},
				Interval{4, 5},
				Interval{6, 7},
				Interval{8, 9},
				Interval{1, 10},
			},
			},
			ans{[]Interval{
				Interval{1, 10},
			},
			},
		},

		question{
			para{[]Interval{
				Interval{1, 4},
				Interval{0, 6},
			},
			},
			ans{[]Interval{
				Interval{0, 6},
			},
			},
		},

		question{
			para{[]Interval{
				Interval{1, 4},
				Interval{2, 5},
			},
			},
			ans{[]Interval{
				Interval{1, 5},
			},
			},
		},

		question{
			para{[]Interval{
				Interval{1, 3},
				Interval{2, 6},
				Interval{8, 10},
				Interval{15, 18},
			}},
			ans{[]Interval{
				Interval{1, 6},
				Interval{8, 10},
				Interval{15, 18},
			},
			},
		},

		question{
			para{[]Interval{
				Interval{1, 6},
			},
			},
			ans{[]Interval{
				Interval{1, 6},
			},
			},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, merge(p.one), "输入:%v", p)
	}
}
