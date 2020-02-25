/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 *
 * https://leetcode-cn.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (46.65%)
 * Likes:    623
 * Dislikes: 0
 * Total Accepted:    32.1K
 * Total Submissions: 68.7K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 * 
 * 
 * 
 * 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢
 * Marcos 贡献此图。
 * 
 * 示例:
 * 
 * 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出: 6
 * 
 */

// @lc code=start
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	var max, maxIndex int
	for i, h := range height {
		if h > max {
			max = h
			maxIndex = i
		}
	}

	var total, leftMax, rightMax int
	for i := 0; i < maxIndex; i++ {
		if leftMax < height[i] {
			leftMax = height[i]
		}
		total += leftMax - height[i]
	}
	for i := len(height)-1; i > maxIndex; i-- {
		if rightMax < height[i] {
			rightMax = height[i]
		}
		total += rightMax - height[i]
	}
	return total
}
// @lc code=end

