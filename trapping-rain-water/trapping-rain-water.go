package leetcode

func bigger(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func smaller(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func trap(height []int) int {
	length := len(height)
	if length <= 2 {
		return 0
	}

	left, right := make([]int, length), make([]int, length)
	left[0], right[length-1] = height[0], height[length-1]

	for i := 1; i < length; i++ {
		left[i] = bigger(left[i-1], height[i])
		right[length-1-i] = bigger(right[length-i], height[length-1-i])
		// left[i] 是height[:i + 1]中的最大值
		// right[i] 是 height[i:] 中的最大值
	}

	water := 0
	for i := 0; i < length; i++ {
		// 凹槽才可以存水
		water += smaller(left[i], right[i]) - height[i]
	}
	return water
}
