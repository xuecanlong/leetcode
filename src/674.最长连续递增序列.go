/*
 * @lc app=leetcode.cn id=674 lang=golang
 *
 * [674] 最长连续递增序列
 *
 * https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence/description/
 *
 * algorithms
 * Easy (42.77%)
 * Likes:    45
 * Dislikes: 0
 * Total Accepted:    12K
 * Total Submissions: 27.9K
 * Testcase Example:  '[1,3,5,4,7]'
 *
 * 给定一个未经排序的整数数组，找到最长且连续的的递增序列。
 * 
 * 示例 1:
 * 
 * 
 * 输入: [1,3,5,4,7]
 * 输出: 3
 * 解释: 最长连续递增序列是 [1,3,5], 长度为3。
 * 尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为5和7在原数组里被4隔开。 
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: [2,2,2,2,2]
 * 输出: 1
 * 解释: 最长连续递增序列是 [2], 长度为1。
 * 
 * 
 * 注意：数组长度不会超过10000。
 * 
 */

// @lc code=start
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max, start, end, left := 0, 0, 0, 0
	for idx, num := range nums {
		if end == 0 || left < num {
			left = num
			end = idx + 1
			if end - start > max {
				max = end - start
			}
		} else {
			left = num
			start = idx
			end = idx + 1
		}
	}
	return max
}
// @lc code=end

