package leetcode

func searchMatrix(mat [][]int, target int) bool {
	m := len(mat)
	if m == 0 {
		return false
	}

	n := len(mat[0])
	if n == 0 {
		return false
	}

	if target < mat[0][0] || target > mat[m-1][n-1] {
		return false
	}

	r := 0
	for r < m && mat[r][0] <= target {
		r++
	}
	r--

	// 二分法寻找 target
	i, j := 0, n-1
	for i <= j {
		med := (i + j) / 2
		switch {
		case mat[r][med] < target:
			i = med + 1
		case target < mat[r][med]:
			j = med - 1
		default:
			return true
		}
	}

	return mat[r][j] == target
}

func searchMatrix2(mat [][]int, target int) bool {
	m := len(mat)
	if m == 0 {
		return false
	}

	n := len(mat[0])
	if n == 0 {
		return false
	}

	if target < mat[0][0] || target > mat[m-1][n-1] {
		return false
	}

	r, first := search2(mat[:m][0], target)

	if first {
		return true
	}
	if search(mat[r], target) == -1 {
		return false
	} else {
		return true
	}
}

// 二分查找法
func search(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		med := (i + j) / 2
		switch {
		case nums[med] < target:
			i = med + 1
		case nums[med] > target:
			j = med - 1
		default:
			return med
		}
	}
	return -1
}

// 二分查找法(返回所在的行数和第一个数是否是target)
func search2(nums []int, target int) (int, bool) {
	i, j := 0, len(nums)-1
	for i <= j {
		med := (i + j) / 2
		switch {
		case nums[med] < target:
			i = med + 1
		case nums[med] > target:
			j = med - 1
		default:
			return med, true
		}
	}
	return j - 1, false
}
