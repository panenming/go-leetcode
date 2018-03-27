## 题目
    Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

    Example:

    Input: "babad"
    Output: "bab"
    Note: "aba" is also a valid answer.

    Example:

    Input: "cbbd"
    Output: "bb"
## 解题思路
    题目要求寻找字符串中的最长回文。
    func isPalindromic(s *string, i, j int ) bool {
        for  i< j {
            if (*s)[i] != (*s)[j] {
                return false
            } 
            i++
            j--
        }
        return true
    }
