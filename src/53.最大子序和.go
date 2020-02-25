/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 *
 * https://leetcode-cn.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (47.70%)
 * Likes:    1294
 * Dislikes: 0
 * Total Accepted:    109.3K
 * Total Submissions: 228.9K
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 * 
 * 示例:
 * 
 * 输入: [-2,1,-3,4,-1,2,1,-5,4],
 * 输出: 6
 * 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
 * 
 * 
 * 进阶:
 * 
 * 如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
 * 
 */

// @lc code=start
func maxSubArray(nums []int) int {
	curMaxSum := 0          // 当前最大连续子序列和
	allMaxSum := nums[0]    // 全局最大连续子序列和
	for _,x := range nums {
		if curMaxSum < 0 { // 当前最大和如果小于0，直接用当前元素值覆盖这个负值
			curMaxSum = x
		} else {        // 如果当前最大和>=0，继续累加
			curMaxSum += x
		}
		if curMaxSum > allMaxSum {
			allMaxSum = curMaxSum
		}
	}
	return allMaxSum
}
// @lc code=end

