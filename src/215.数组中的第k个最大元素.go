/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 *
 * https://leetcode-cn.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (58.75%)
 * Likes:    280
 * Dislikes: 0
 * Total Accepted:    51.5K
 * Total Submissions: 87.7K
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * 在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
 * 
 * 示例 1:
 * 
 * 输入: [3,2,1,5,6,4] 和 k = 2
 * 输出: 5
 * 
 * 
 * 示例 2:
 * 
 * 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
 * 输出: 4
 * 
 * 说明: 
 * 
 * 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
 * 
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	return doFindKthLargest(nums, k, 0, len(nums)-1)
}

func doFindKthLargest(nums []int, k int, start, end int) int {
	nlen := len(nums)
	targetPos := nlen - k

	for {
		pivotIndex := partition(nums, start, end)
		// 比较这个数所在的位置跟目标位置的大小,小则目标在后面部分故排序后面部分,大则目标在前面部分故排序前面部分
		if pivotIndex == targetPos {
			return nums[pivotIndex]
		} else if pivotIndex < targetPos {
			start = pivotIndex + 1
		} else {
			end = pivotIndex - 1
		}
	}
}

// 做一次不完整的快排,得到一个数在数组中的位置
func partition(nums []int, start, end int) int {
	pivot := nums[start]
	left, right := start+1, end

	for left <= right {
		for left <= right && nums[left] <= pivot {
			left++
		}
		for left <= right && nums[right] > pivot {
			right--
		}
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	nums[right], nums[start] = nums[start], nums[right]
	return right
}
// @lc code=end

