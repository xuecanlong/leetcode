/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] 阶乘后的零
 *
 * https://leetcode-cn.com/problems/factorial-trailing-zeroes/description/
 *
 * algorithms
 * Easy (39.59%)
 * Likes:    159
 * Dislikes: 0
 * Total Accepted:    20.6K
 * Total Submissions: 52K
 * Testcase Example:  '3'
 *
 * 给定一个整数 n，返回 n! 结果尾数中零的数量。
 * 
 * 示例 1:
 * 
 * 输入: 3
 * 输出: 0
 * 解释: 3! = 6, 尾数中没有零。
 * 
 * 示例 2:
 * 
 * 输入: 5
 * 输出: 1
 * 解释: 5! = 120, 尾数中有 1 个零.
 * 
 * 说明: 你算法的时间复杂度应为 O(log n) 。
 * 
 */

// @lc code=start
func trailingZeroes(n int) int {
	//read more why we count the 5x, 25x, 125x... : https://brilliant.org/wiki/trailing-number-of-zeros/
	var i, res int = 5, 0

	for n >= i {
		res += n / i
		i *= 5
	}

	return res
}
// @lc code=end

