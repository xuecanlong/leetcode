/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 翻转字符串里的单词
 *
 * https://leetcode-cn.com/problems/reverse-words-in-a-string/description/
 *
 * algorithms
 * Medium (31.98%)
 * Likes:    82
 * Dislikes: 0
 * Total Accepted:    20.1K
 * Total Submissions: 62.9K
 * Testcase Example:  '"the sky is blue"'
 *
 * 给定一个字符串，逐个翻转字符串中的每个单词。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入: "the sky is blue"
 * 输出: "blue is sky the"
 * 
 * 
 * 示例 2：
 * 
 * 输入: "  hello world!  "
 * 输出: "world! hello"
 * 解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
 * 
 * 
 * 示例 3：
 * 
 * 输入: "a good   example"
 * 输出: "example good a"
 * 解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
 * 
 * 
 * 
 * 
 * 说明：
 * 
 * 
 * 无空格字符构成一个单词。
 * 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
 * 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
 * 
 * 
 * 
 * 
 * 进阶：
 * 
 * 请选用 C 语言的用户尝试使用 O(1) 额外空间复杂度的原地解法。
 * 
 */

// @lc code=start

func reverseWords(s string) string {
	// 去除首尾空格
	s = strings.TrimSpace(s)
	// 整个字符串翻转
	s = reverse(s, 0, len(s)-1)
	start := 0
	for {
		if i := strings.Index(s[start:], " "); i == -1 {
			break
		} else {
			if i == 0 {
				// 去除多余空格
				s = trimSpace(s, start)
				start = start
			} else {
				// 单词翻转
				s = reverse(s, start, start+i-1)
				start = start + i + 1
			}
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

func trimSpace(s string, delete int) string {
	sb := []byte(s)
	sb = append(sb[:delete], sb[delete + 1:]...)
	return string(sb)
}
// @lc code=end

