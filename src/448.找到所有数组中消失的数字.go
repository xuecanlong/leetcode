/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 *
 * https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/description/
 *
 * algorithms
 * Easy (53.54%)
 * Likes:    220
 * Dislikes: 0
 * Total Accepted:    17.2K
 * Total Submissions: 31.9K
 * Testcase Example:  '[4,3,2,7,8,2,3,1]'
 *
 * 给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。
 * 
 * 找到所有在 [1, n] 范围之间没有出现在数组中的数字。
 * 
 * 您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。
 * 
 * 示例:
 * 
 * 
 * 输入:
 * [4,3,2,7,8,2,3,1]
 * 
 * 输出:
 * [5,6]
 * 
 * 
 */

// @lc code=start

// # 将所有正数作为数组下标，置对应数组值为负值。那么，仍为正数的位置即为（未出现过）消失的数字。
// # 举个例子：
// # 原始数组：[4,3,2,7,8,2,3,1]
// # 重置后为：[-4,-3,-2,-7,8,2,-3,-1]
// # 结论：[8,2] 分别对应的index为[5,6]（消失的数字）
func myAbs(a int) int {
	if a > 0 {
		return a
	} else {
		return a * -1
	}
}


func findDisappearedNumbers(nums []int) []int {
	ret := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[myAbs(nums[i])-1] > 0 {
			nums[myAbs(nums[i])-1] *= -1
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			ret = append(ret, i+1)
		}
	}
	return ret
}
// @lc code=end

