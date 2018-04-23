package leetcode

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	// 用栈存放数字
	nums := make([]int, 0, len(tokens))

	for _, s := range tokens {
		if s == "+" ||
			s == "-" ||
			s == "*" ||
			s == "/" {
			b, a := nums[len(nums)-1], nums[len(nums)-2]
			nums = nums[:len(nums)-2]
			nums = append(nums, compute(a, b, s))
		} else {
			temp, _ := strconv.Atoi(s)
			nums = append(nums, temp)
		}
	}
	return nums[0]
}

// 计算
func compute(a, b int, opt string) int {
	switch opt {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	default:
		return a / b
	}
}
