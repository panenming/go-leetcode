package leetcode

// 算法有误，因为s1和s2中的相同值会导致数据有误
func isInterleave2(s1, s2, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if l1 == 0 || l2 == 0 || l3 == 0 {
		return false
	}
	if l1+l2 != l3 {
		return false
	}

	index1, index2 := 0, 0
	for i := 0; i < l3; i++ {
		if index1 < l1 && s3[i] == s1[index1] {
			index1++
		} else if index2 < l2 && s3[i] == s2[index2] {
			index2++
		} else {
			return false
		}
	}
	return true

}

func isInterleave(s1, s2, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}

	// dp[i][j][i+j] == true 表示 s1[:i] 和 s2[:j] 可以生成 s3[:i+j]
	dp := make([][][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]bool, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([]bool, m+n+1)
		}
	}
	dp[0][0][0] = true

	for i := 1; i <= m; i++ {
		dp[i][0][i] = dp[i-1][0][i-1] && s1[i-1] == s3[i-1]
	}

	for j := 1; j <= n; j++ {
		dp[0][j][j] = dp[0][j-1][j-1] && s2[j-1] == s3[j-1]
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j][i+j] = (dp[i-1][j][i+j-1] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1][i+j-1] && s2[j-1] == s3[i+j-1])
		}
	}

	return dp[m][n][m+n]
}
