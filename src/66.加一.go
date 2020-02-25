/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 *
 * https://leetcode-cn.com/problems/plus-one/description/
 *
 * algorithms
 * Easy (41.25%)
 * Likes:    365
 * Dislikes: 0
 * Total Accepted:    87.8K
 * Total Submissions: 211.6K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
 * 
 * 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
 * 
 * 你可以假设除了整数 0 之外，这个整数不会以零开头。
 * 
 * 示例 1:
 * 
 * 输入: [1,2,3]
 * 输出: [1,2,4]
 * 解释: 输入数组表示数字 123。
 * 
 * 
 * 示例 2:
 * 
 * 输入: [4,3,2,1]
 * 输出: [4,3,2,2]
 * 解释: 输入数组表示数字 4321。
 * 
 * 
 */

// @lc code=start
func plusOne(digits []int) []int {
	add := 1
	for i:= len(digits) - 1; i >= 0; i -- {
		if add == 0 {
			break
		}
		digits[i] = digits[i] + add
		add = digits[i] / 10
		digits[i] = digits[i] % 10
	}
	for add > 0 {
		digits = append([]int{add % 10}, digits...)
		add = add / 10
	}
	return digits
}
// @lc code=end

