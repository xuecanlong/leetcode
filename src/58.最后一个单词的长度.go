/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 *
 * https://leetcode-cn.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (31.29%)
 * Likes:    137
 * Dislikes: 0
 * Total Accepted:    48.7K
 * Total Submissions: 155K
 * Testcase Example:  '"Hello World"'
 *
 * 给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
 * 
 * 如果不存在最后一个单词，请返回 0 。
 * 
 * 说明：一个单词是指由字母组成，但不包含任何空格的字符串。
 * 
 * 示例:
 * 
 * 输入: "Hello World"
 * 输出: 5
 * 
 * 
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	for s[len(s)-1] == ' ' { //去掉最后的空格
		s = s[0 : len(s)-1]
		if len(s) == 0 {
			return 0
		}
	}

	length := 0
	for s[len(s)-1] != ' ' { //从后往前判断
		s = s[0 : len(s)-1]
		length ++
		if len(s) == 0 {
			break
		}
	}
	return length
}

// @lc code=end

