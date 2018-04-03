package leetcode

func generateParenthesis(n int) []string {
	res := []string{}
	gen(n, n, "", &res)
	return res
}

func gen(left, right int, substr string, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, substr)
		return
	}
	if left > 0 {
		gen(left-1, right, substr+"(", res)
	}
	if right > 0 && left < right {
		gen(left, right-1, substr+")", res)
	}
}
