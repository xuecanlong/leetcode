/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 *
 * https://leetcode-cn.com/problems/longest-consecutive-sequence/description/
 *
 * algorithms
 * Hard (46.88%)
 * Likes:    172
 * Dislikes: 0
 * Total Accepted:    16.9K
 * Total Submissions: 36.1K
 * Testcase Example:  '[100,4,200,1,3,2]'
 *
 * 给定一个未排序的整数数组，找出最长连续序列的长度。
 * 
 * 要求算法的时间复杂度为 O(n)。
 * 
 * 示例:
 * 
 * 输入: [100, 4, 200, 1, 3, 2]
 * 输出: 4
 * 解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
 * 
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	max := 0
	for i, _ := range m {
		tmp := helper(m, i)
		if max < tmp {
			max = tmp
		}
	}
	return max
}

//查找相邻数字是否存在
func helper(m map[int]bool, index int) int {
	if !m[index] {
		return 0
	}
	m[index] = false
	return helper(m, index + 1) + helper(m, index - 1) + 1
}
// @lc code=end

