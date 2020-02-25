/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 *
 * https://leetcode-cn.com/problems/merge-sorted-array/description/
 *
 * algorithms
 * Easy (45.51%)
 * Likes:    332
 * Dislikes: 0
 * Total Accepted:    79.3K
 * Total Submissions: 173.8K
 * Testcase Example:  '[1,2,3,0,0,0]\n3\n[2,5,6]\n3'
 *
 * 给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
 * 
 * 说明:
 * 
 * 
 * 初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
 * 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
 * 
 * 
 * 示例:
 * 
 * 输入:
 * nums1 = [1,2,3,0,0,0], m = 3
 * nums2 = [2,5,6],       n = 3
 * 
 * 输出: [1,2,2,3,5,6]
 * 
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) []int {
	for m > 0 || n > 0{
		if m == 0 {
			nums1[m+n-1] = nums2[n -1]
			n--
			continue
		}
		if n == 0 {
			nums1[m+n-1] = nums1[m -1]
			m--
			continue
		}
		if nums1[m -1] > nums2[n -1] {
			nums1[m+n-1] = nums1[m -1]
			m--
		} else {
			nums1[m+n-1] = nums2[n -1]
			n--
		}
	}
	return nums1
}
// @lc code=end

