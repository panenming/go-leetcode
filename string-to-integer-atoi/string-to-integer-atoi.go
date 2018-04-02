package leetcode

import (
	"math"
	"strings"
)

func atoi(str string) int {
	s := strings.TrimSpace(str)
	if len(s) == 0 {
		return 0
	}

	// 取出 s 的符号和主体 x
	sign, x := getSign(s)

	// 裁剪x丢弃混入的非数字字符
	x = trim(x)

	// 根据sigh和x，转换成int
	return convert(sign, x)

}

func getSign(s string) (int, string) {
	sign := 1
	switch s[0] {
	case '-':
		s = s[1:]
		sign = -1.0
	case '+':
		s = s[1:]
	default:
	}
	return sign, s
}

func trim(s string) string {
	for i := range s {
		if s[i] < '0' || s[i] > '9' {
			return s[:i]
		}
	}
	return s
}

func convert(sign int, s string) int {
	base := 1 * sign
	res := 0
	yes := false
	for i := len(s) - 1; i >= 0; i-- {
		// 48 asscii to int
		res, yes = isOverflow(res + (int(s[i])-48)*base)
		if yes {
			return res
		}
		base *= 10
	}
	return res
}

func isOverflow(i int) (int, bool) {
	switch {
	case i > math.MaxInt32:
		return math.MaxInt32, true
	case i < math.MinInt32:
		return math.MinInt32, true
	default:
		return i, false
	}
}
