/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 *
 * https://leetcode-cn.com/problems/maximal-square/description/
 *
 * algorithms
 * Medium (38.28%)
 * Likes:    147
 * Dislikes: 0
 * Total Accepted:    12.9K
 * Total Submissions: 33.8K
 * Testcase Example:  '[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]'
 *
 * 在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。
 * 
 * 示例:
 * 
 * 输入: 
 * 
 * 1 0 1 0 0
 * 1 0 1 1 1
 * 1 1 1 1 1
 * 1 0 0 1 0
 * 
 * 输出: 4
 * 
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	rows := len(matrix)
	cols := 0
	if rows > 0 {
		cols = len(matrix[0])
	}
	dp := make([]int, cols + 1)
	maxsqlen, prev := 0, 0
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			temp := dp[j];
			if matrix[i - 1][j - 1] == '1' {
				// 求prev, dp[j - 1], dp[j] 中的最小值 + 1
				min1 := 0
				if dp[j - 1] > prev {
					min1 = prev
				} else {
					min1 = dp[j - 1]
				}
				if dp[j] > min1 {
					dp[j] = min1 + 1
				} else {
					dp[j] = dp[j] + 1
				}

				// 计算最大值
				if dp[j] > maxsqlen {
					maxsqlen = dp[j]
				}
			} else {
				dp[j] = 0;
			}
			prev = temp
		}
	}
	return maxsqlen * maxsqlen
}
// @lc code=end

