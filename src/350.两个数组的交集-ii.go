/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 *
 * https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/description/
 *
 * algorithms
 * Easy (44.64%)
 * Likes:    206
 * Dislikes: 0
 * Total Accepted:    54.9K
 * Total Submissions: 121.6K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * 给定两个数组，编写一个函数来计算它们的交集。
 * 
 * 示例 1:
 * 
 * 输入: nums1 = [1,2,2,1], nums2 = [2,2]
 * 输出: [2,2]
 * 
 * 
 * 示例 2:
 * 
 * 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * 输出: [4,9]
 * 
 * 说明：
 * 
 * 
 * 输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
 * 我们可以不考虑输出结果的顺序。
 * 
 * 
 * 进阶:
 * 
 * 
 * 如果给定的数组已经排好序呢？你将如何优化你的算法？
 * 如果 nums1 的大小比 nums2 小很多，哪种方法更优？
 * 如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
 * 
 * 
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	for _, num := range nums1 {
		if count, ok := m[num]; !ok {
			m[num] = 1
		} else {
			m[num] = count + 1
		}
	}
	var res []int
	for _, num := range nums2 {
		if count, ok := m[num]; ok && count > 0 {
			m[num] = count - 1
			res = append(res, num)
		}
	}
	return res
}

// Q1:
// 如果给定的数组已经排好序呢？你将如何优化你的算法？
// A1:
// nums1和nums2的长度分别是m和n，假设m<=n
// 用计数器count1记下nums1中元素的个数，然后对每个元素进行遍历，用二分查找找到该数在nums2中的左右边界，例如2在[1,2,2,2,4]中的左右边界为[1, 3]。复杂度为O(mlogn)

// Q2:
// 如果 nums1 的大小比 nums2 小很多，哪种方法更优？
// A2:
// 因为双指针遍历的时间复杂度为O(m+n)，所以当n>>m时，二分查找优于双指针遍历

// @lc code=end

