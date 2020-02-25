/*
 * @lc app=leetcode.cn id=1037 lang=golang
 *
 * [1037] 有效的回旋镖
 *
 * https://leetcode-cn.com/problems/valid-boomerang/description/
 *
 * algorithms
 * Easy (37.97%)
 * Likes:    7
 * Dislikes: 0
 * Total Accepted:    2.4K
 * Total Submissions: 6.4K
 * Testcase Example:  '[[1,1],[2,3],[3,2]]'
 *
 * 回旋镖定义为一组三个点，这些点各不相同且不在一条直线上。
 * 
 * 给出平面上三个点组成的列表，判断这些点是否可以构成回旋镖。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：[[1,1],[2,3],[3,2]]
 * 输出：true
 * 
 * 
 * 示例 2：
 * 
 * 输入：[[1,1],[2,2],[3,3]]
 * 输出：false
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * points.length == 3
 * points[i].length == 2
 * 0 <= points[i][j] <= 100
 * 
 * 
 */

// @lc code=start
func isBoomerang(points [][]int) bool {
	dx1 := points[1][0] - points[0][0]
	dx2 := points[1][0] - points[2][0]
	dy1 := points[1][1] - points[0][1]
	dy2 := points[1][1] - points[2][1]
	
	return dx1*dy2 != dx2*dy1
}
// @lc code=end

