package leetcode

func maxArea(height []int) int {
	// 从两端开始寻找，至少保证了宽度是最大值
	i, j := 0, len(height)-1
	max := 0
	for i < j {
		a, b := height[i], height[j]
		h := min(a, b)
		area := h * (j - i)
		if max < area {
			max = area
		}

		//短板在哪边移动哪边
		if a < b {
			i++
		} else {
			j--
		}
	}
	return max
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}
