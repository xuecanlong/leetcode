/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个排序链表
 *
 * https://leetcode-cn.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (47.57%)
 * Likes:    360
 * Dislikes: 0
 * Total Accepted:    49.5K
 * Total Submissions: 103.9K
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * 合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
 * 
 * 示例:
 * 
 * 输入:
 * [
 * 1->4->5,
 * 1->3->4,
 * 2->6
 * ]
 * 输出: 1->1->2->3->4->4->5->6
 * 
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}
	length := len(lists)
	if length == 0 {
		return nil
	}
	for length > 1 {
		for i := 0; i < length / 2; i++ {
			lists[i] = mergeTwoLists(lists[i], lists[length - i - 1])
			lists = lists[:length - i - 1]
		}
		length = len(lists)
	}
	return lists[0]
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res
	for l1 != nil || l2 != nil {
		if l1 == nil {
			cur.Next = l2
			l2 = nil
			continue
		}
		if l2 == nil {
			cur.Next = l1
			l1 = nil
			continue
		}

		if l1.Val > l2.Val {
			cur.Next = l2
			cur, l2 = cur.Next, l2.Next
		} else {
			cur.Next = l1
			cur, l1 = cur.Next, l1.Next
		}
	}
	return res.Next
}

// @lc code=end

