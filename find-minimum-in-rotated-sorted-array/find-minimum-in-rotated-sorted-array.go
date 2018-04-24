package leetcode

func findMin(a []int) int {
	len := len(a)
	i := 1
	for i < len && a[i-1] < a[i] {
		i++
	}
	return a[i%len]
}

// 存在重复值
func findMinDups(a []int) int {
	l := len(a)
	if a[0] < a[l-1] {
		return a[0]
	}

	i, j := 0, l-1
	for i < j {
		if a[i] > a[i+1] {
			return a[i+1]
		}

		if a[j-1] > a[j] {
			return a[j]
		}

		i++
		j--
	}
	return a[i]
}
