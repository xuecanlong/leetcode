/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (31.72%)
 * Likes:    2530
 * Dislikes: 0
 * Total Accepted:    240.2K
 * Total Submissions: 757.2K
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 * 
 * 示例 1:
 * 
 * 输入: "abcabcbb"
 * 输出: 3 
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 * 
 * 
 * 示例 2:
 * 
 * 输入: "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 * 
 * 
 * 示例 3:
 * 
 * 输入: "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 * 
 * 
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	m, max, left := make(map[rune]int), 0, 0
	for idx, c := range s {
		if data, okay := m[c]; okay == true && data >= left {
			if idx - left > max {
				max = idx - left
			}
			left = m[c] + 1
		}
		m[c] = idx
	}
	if len(s) - left > max {
		max = len(s) - left
	}
	return max
}
// @lc code=end

