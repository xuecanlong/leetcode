/*
 * @lc app=leetcode.cn id=453 lang=golang
 *
 * [453] 最小移动次数使数组元素相等
 *
 * https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements/description/
 *
 * algorithms
 * Easy (52.55%)
 * Likes:    77
 * Dislikes: 0
 * Total Accepted:    6.1K
 * Total Submissions: 11.5K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个长度为 n 的非空整数数组，找到让数组所有元素相等的最小移动次数。每次移动可以使 n - 1 个元素增加 1。
 * 
 * 示例:
 * 
 * 
 * 输入:
 * [1,2,3]
 * 
 * 输出:
 * 3
 * 
 * 解释:
 * 只需要3次移动（注意每次移动会增加两个元素的值）：
 * 
 * [1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
 * 
 * 
 */

// @lc code=start
func minMoves(nums []int) int {
	length := len(nums)
	if length < 2 {
		return 0
	}
	min := 0
	res := 0
	for i := 0; i < length; i++ {
		if i == 0 {
			min = nums[i]
		} else {
			if nums[i] < min {
				min = nums[i]
			}
		}
	}
	for _, num := range nums {
		res += num - min
	}
	return res
}
// @lc code=end

