/*
 * @lc app=leetcode.cn id=611 lang=golang
 *
 * [611] 有效三角形的个数
 *
 * https://leetcode-cn.com/problems/valid-triangle-number/description/
 *
 * algorithms
 * Medium (46.47%)
 * Likes:    43
 * Dislikes: 0
 * Total Accepted:    3K
 * Total Submissions: 6.4K
 * Testcase Example:  '[2,2,3,4]'
 *
 * 给定一个包含非负整数的数组，你的任务是统计其中可以组成三角形三条边的三元组个数。
 * 
 * 示例 1:
 * 
 * 
 * 输入: [2,2,3,4]
 * 输出: 3
 * 解释:
 * 有效的组合是: 
 * 2,3,4 (使用第一个 2)
 * 2,3,4 (使用第二个 2)
 * 2,2,3
 * 
 * 
 * 注意:
 * 
 * 
 * 数组长度不超过1000。
 * 数组里整数的范围为 [0, 1000]。
 * 
 * 
 */

// @lc code=start
// 首先对数组排序。
// 固定最长的一条边，运用双指针扫描
// 如果 nums[l] + nums[r] > nums[i]，同时说明 nums[l + 1] + nums[r] > nums[i], ..., nums[r - 1] + nums[r] > nums[i]，满足的条件的有 r - l 种，r 左移进入下一轮。
// 如果 nums[l] + nums[r] <= nums[i]，l 右移进入下一轮。
// 枚举结束后，总和就是答案。
// 时间复杂度为 O(n^2)
func triangleNumber(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	n := len(nums)
	res := 0
	for i := n -1; i >= 2; i-- {
		l := 0
		r := i - 1
		for l < r {
			if nums[l] + nums[r] > nums[i] {
				res += r - l
				r--
			} else {
				l++
			}
		}
	}
	return res
}
// @lc code=end

