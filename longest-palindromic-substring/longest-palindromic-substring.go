package leetcode

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	// 最长回文的首字符索引，和最长回文的长度
	begin, maxLen := 0, 1
	// 在 for 循环中
	// b 代表回文的**首**字符索引号，
	// e 代表回文的**尾**字符索引号，
	// i 代表回文`正中间段`首字符的索引号
	// 在每一次for循环中
	// 先从i开始，利用`正中间段`所有字符相同的特性，让b，e分别指向`正中间段`的首尾
	// 再从`正中间段`向两边扩张，让b，e分别指向此`正中间段`为中心的最长回文的首尾
	for i := 0; i < len(s); {
		if len(s)-i <= maxLen/2 {
			break
		}
		b, e := i, i
		for e < len(s)-1 && s[e+1] == s[e] {
			e++
		}
		i = e + 1

		for e < len(s)-1 && b > 0 && s[e+1] == s[b-1] {
			e++
			b--
			// 循环结束后，s[b:e+1]是这次能找到的最长回文。
		}

		newLen := e + 1 - b
		// 创新记录的话，就更新记录
		if newLen > maxLen {
			begin = b
			maxLen = newLen
		}
	}
	return s[begin : begin+maxLen]
}
