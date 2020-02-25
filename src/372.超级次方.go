/*
 * @lc app=leetcode.cn id=372 lang=golang
 *
 * [372] 超级次方
 *
 * https://leetcode-cn.com/problems/super-pow/description/
 *
 * algorithms
 * Medium (36.19%)
 * Likes:    43
 * Dislikes: 0
 * Total Accepted:    2K
 * Total Submissions: 5.6K
 * Testcase Example:  '2\n[3]'
 *
 * 你的任务是计算 a^b 对 1337 取模，a 是一个正整数，b 是一个非常大的正整数且会以数组形式给出。
 * 
 * 示例 1:
 * 
 * 输入: a = 2, b = [3]
 * 输出: 8
 * 
 * 
 * 示例 2:
 * 
 * 输入: a = 2, b = [1,0]
 * 输出: 1024
 * 
 */

// @lc code=start
func superPow(a int, b []int) int {
	var pow func(x, y int) int
	pow = func(x, y int) int {
		switch y {
		case 0:
			return 1
		case 1:
			return x % 1337
		default:
			return pow(x%1337, y/2) * pow(x%1337, y-y/2) % 1337
		}
	}
	out := 1
	for _, e := range b {
		out = pow(out, 10) * pow(a, e) % 1337
	}
	return out
}
// @lc code=end

