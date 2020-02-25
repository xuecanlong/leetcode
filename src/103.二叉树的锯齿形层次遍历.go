/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层次遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/description/
 *
 * algorithms
 * Medium (52.23%)
 * Likes:    103
 * Dislikes: 0
 * Total Accepted:    20.5K
 * Total Submissions: 39.2K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
 * 
 * 例如：
 * 给定二叉树 [3,9,20,null,null,15,7],
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 
 * 返回锯齿形层次遍历如下：
 * 
 * [
 * ⁠ [3],
 * ⁠ [20,9],
 * ⁠ [15,7]
 * ]
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
 func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	var res [][]int
	leftToRight := false
	for {   // 一层一层取出链表
		qlen := len(queue)
		if qlen == 0 {
			break
		}

		var arr []int
		// 将这层链表放入数组
		for i := qlen - 1; i >= 0; i-- {
			node := queue[0]
			arr = append(arr, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}

		if leftToRight { //为了锯齿层次倒序
			reverse(arr)
		}
		res = append(res, arr)
		leftToRight = !leftToRight
	}
	return res
}

func reverse(list []int) {
	if len(list) > 0 {
		for i := 0; i < len(list)/2; i++ {
			list[i], list[len(list)-1-i] = list[len(list)-1-i], list[i]
		}
	}
}

// @lc code=end

