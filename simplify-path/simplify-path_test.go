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
	one string
}

type ans struct {
	one string
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		question{
			para{
				"/home/",
			},
			ans{
				"/home",
			},
		},

		question{
			para{
				"/a/./b/../../c/",
			},
			ans{
				"/c",
			},
		},

		question{
			para{
				"/../",
			},
			ans{
				"/",
			},
		},

		question{
			para{
				"/home//foo/",
			},
			ans{
				"/home/foo",
			},
		},

		question{
			para{
				"/home/./foo/",
			},
			ans{
				"/home/foo",
			},
		},

		question{
			para{
				"/home////////////////////////foo/",
			},
			ans{
				"/home/foo",
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("---")
		ast.Equal(a.one, simplifyPath(p.one), "输入:%v", p)
	}
}
