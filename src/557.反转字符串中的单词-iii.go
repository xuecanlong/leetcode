/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 *
 * https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/description/
 *
 * algorithms
 * Easy (67.96%)
 * Likes:    123
 * Dislikes: 0
 * Total Accepted:    30.5K
 * Total Submissions: 44.9K
 * Testcase Example:  `"Let's take LeetCode contest"`
 *
 * 给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
 * 
 * 示例 1:
 * 
 * 
 * 输入: "Let's take LeetCode contest"
 * 输出: "s'teL ekat edoCteeL tsetnoc" 
 * 
 * 
 * 注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
 * 
 */

// @lc code=start
func reverseWords(s string) string {
	start := 0
	for {
		if i := strings.Index(s[start:], " "); i == -1 {
			break
		} else {
			s = reverse(s, start, start+i-1)
			start = start + i + 1
		}
	}

	if start < len(s) {
		s = reverse(s, start, len(s)-1)
	}
	return s
}

func reverse(s string, start, end int) string {
	sb := []byte(s)
	for start < end {
		sb[start], sb[end] = sb[end], sb[start]
		start, end = start+1, end-1
	}
	return string(sb)
}
// @lc code=end

