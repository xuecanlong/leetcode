/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
 *
 * https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/description/
 *
 * algorithms
 * Easy (67.22%)
 * Likes:    259
 * Dislikes: 0
 * Total Accepted:    34.9K
 * Total Submissions: 51.6K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * 将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
 * 
 * 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
 * 
 * 示例:
 * 
 * 给定有序数组: [-10,-3,0,5,9],
 * 
 * 一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
 * 
 * ⁠     0
 * ⁠    / \
 * ⁠  -3   9
 * ⁠  /   /
 * ⁠-10  5
 * 
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil {
		return nil;
	}
	return helper(nums, 0, len(nums) - 1)
}

func helper (nums []int, low int, high int)  *TreeNode {
	if (low > high) { // Done
		return nil
	}
	mid := low + (high - low) / 2
	node := new(TreeNode)
	node.Val = nums[mid]
	node.Left = helper(nums, low, mid - 1)
	node.Right = helper(nums, mid + 1, high)
	return node
}
// @lc code=end

