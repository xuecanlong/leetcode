/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 *
 * https://leetcode-cn.com/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (64.27%)
 * Likes:    194
 * Dislikes: 0
 * Total Accepted:    42.7K
 * Total Submissions: 66.2K
 * Testcase Example:  '5'
 *
 * 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
 * 
 * 
 * 
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 * 
 * 示例:
 * 
 * 输入: 5
 * 输出:
 * [
 * ⁠    [1],
 * ⁠   [1,1],
 * ⁠  [1,2,1],
 * ⁠ [1,3,3,1],
 * ⁠[1,4,6,4,1]
 * ]
 * 
 */

// @lc code=start
func generate(numRows int) [][]int {
	res := [][]int{}
	for i := 0; i < numRows; i++ {
		res = append(res, []int{})
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				res[i] = append(res[i], 1)
				continue
			}
			res[i] = append(res[i], res[i - 1][j - 1] + res[i - 1][j])
		}
	}
	return res
}
// @lc code=end

