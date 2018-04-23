package leetcode

func candy(ratings []int) int {
	n := len(ratings)
	if n <= 1 {
		return n
	}

	left := make([]int, n)
	right := make([]int, n)
	left[0] = 1
	right[n-1] = 1
	for i := 1; i < n; i++ {
		if ratings[i-1] < ratings[i] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}

		if ratings[n-i-1] > ratings[n-i] {
			// n-i-1 比右边的大
			// 所以，他的数量要比右边的多一个
			right[n-i-1] = right[n-i] + 1
		} else {
			right[n-i-1] = 1
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		res += max(left[i], right[i])
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
