/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 *
 * https://leetcode-cn.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (27.43%)
 * Likes:    1405
 * Dislikes: 0
 * Total Accepted:    131.3K
 * Total Submissions: 477.4K
 * Testcase Example:  '"babad"'
 *
 * 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
 * 
 * 示例 1：
 * 
 * 输入: "babad"
 * 输出: "bab"
 * 注意: "aba" 也是一个有效答案。
 * 
 * 
 * 示例 2：
 * 
 * 输入: "cbbd"
 * 输出: "bb"
 * 
 * 
 */

// @lc code=start
func longestPalindrome(s string) string {
	min := 0
	max := 0
	slen := len(s)

	for i := 0; i < slen; i++ {
	    // find longest palindromic substring with odd length and center at i
		l := i
		h := i
		for l >= 0 && h < slen && s[l] == s[h] {
			l--
			h++
		}
		if h-l-1 > max-min {
			min = l + 1
			max = h
		}

	    // find longest palindromic substring with even length and center at i,i+1
		l = i
		h = i + 1
		for l >= 0 && h < slen && s[l] == s[h] {
			l--
			h++
		}
		if h-l-1 > max-min {
			min = l + 1
			max = h
		}
	}

	return s[min:max]
}
// @lc code=end

