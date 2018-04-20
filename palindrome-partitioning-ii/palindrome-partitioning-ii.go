package leetcode

func minCut(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = i - 1
	}

	for i := 0; i < n+1; i++ {
		for j := 0; 0 <= i-j && i+j < n && s[i-j] == s[i+j]; j++ {
			dp[i+j+1] = min(dp[i-j]+1, dp[i+j+1])
		}

		for j := 1; 0 <= i-j+1 && i+j < n && s[i-j+1] == s[i+j]; j++ {
			dp[i+j+1] = min(dp[i-j+1]+1, dp[i+j+1])
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
