package leetcode

func numDecdings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)

	dp[0], dp[1] = 1, one(s[0:1])

	for i := 2; i <= n; i++ {
		w1, w2 := one(s[i-1:i]), two(s[i-2:i])
		dp[i] = dp[i-1]*w1 + dp[i-2]*w2
		if dp[i] == 0 {
			// 子字符串 s[:i+1] 的组成方式数为 0
			// 则，s 的组成方式数肯定也为 0
			// 这时可以提前结束
			return 0
		}
	}

	return dp[n]

}

// 检查 s 是否为合格的单字符
// len(s) == 1
// 合格，  返回 1
// 不合格， 返回 0
// 注意：
//     题目保证了 s 只含有数字
func one(s string) int {
	if s == "0" {
		return 0
	}
	return 1
}

// 检查 s 是否为合格的双字符
// len(s) == 2
// 合格，  返回 1
// 不合格， 返回 0
// 注意：
//     题目保证了 s 只含有数字
func two(s string) int {
	switch s[0] {
	case '1':
		return 1
	case '2':
		if s[1] == '7' || s[1] == '8' || s[1] == '9' {
			return 0
		}
		return 1
	default:
		// '3' <= s[0] && s[0] <= '9'，或
		// s[0] == '0'
		return 0
	}
}
