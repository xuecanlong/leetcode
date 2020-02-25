/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 *
 * https://leetcode-cn.com/problems/longest-common-subsequence/description/
 *
 * algorithms
 * Medium (54.67%)
 * Likes:    8
 * Dislikes: 0
 * Total Accepted:    1.7K
 * Total Submissions: 2.9K
 * Testcase Example:  '"abcde"\n"ace"'
 *
 * 给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列。
 * 
 * 一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
 * 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde"
 * 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。
 * 
 * 若这两个字符串没有公共子序列，则返回 0。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入：text1 = "abcde", text2 = "ace" 
 * 输出：3  
 * 解释：最长公共子序列是 "ace"，它的长度为 3。
 * 
 * 
 * 示例 2:
 * 
 * 输入：text1 = "abc", text2 = "abc"
 * 输出：3
 * 解释：最长公共子序列是 "abc"，它的长度为 3。
 * 
 * 
 * 示例 3:
 * 
 * 输入：text1 = "abc", text2 = "def"
 * 输出：0
 * 解释：两个字符串没有公共子序列，返回 0。
 * 
 * 
 * 
 * 
 * 提示:
 * 
 * 
 * 1 <= text1.length <= 1000
 * 1 <= text2.length <= 1000
 * 输入的字符串只含有小写英文字符。
 * 
 * 
 */

// @lc code=start

func longestCommonSubsequence(text1 string, text2 string) int {
	len1 := len(text1)
	len2 := len(text2)
	if len1 == 0 || len2 == 0 {
		return 0
	}
	cols, rows := 0, 0
	colText, rowText := "", ""
	if len1 > len2 {
		cols = len2
		colText = text2
		rows = len1
		rowText = text1
	} else {
		cols = len1
		colText = text1
		rows = len2
		rowText = text2
	}
	dp := make([]int, cols + 1)
	for i := 1; i <= rows; i++ {
		tmp := 0
		for j := 1; j <= cols; j++ {
			prev := tmp
			tmp = dp[j]
			if rowText[i - 1] == colText[j - 1] {
				dp[j] = prev + 1
			} else {
				if dp[j - 1] > dp[j] {
					dp[j] = dp[j - 1]
				}
			}
		}
	}
	return dp[cols]
}
// @lc code=end

