/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 *
 * https://leetcode-cn.com/problems/sort-list/description/
 *
 * algorithms
 * Medium (62.26%)
 * Likes:    304
 * Dislikes: 0
 * Total Accepted:    26.9K
 * Total Submissions: 43.1K
 * Testcase Example:  '[4,2,1,3]'
 *
 * 在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
 * 
 * 示例 1:
 * 
 * 输入: 4->2->1->3
 * 输出: 1->2->3->4
 * 
 * 
 * 示例 2:
 * 
 * 输入: -1->5->3->4->0
 * 输出: -1->0->3->4->5
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
 func sortList(head *ListNode) *ListNode {
	// 如果 head为空或者head就一位,直接返回
	if head == nil || head.Next == nil {
		return head
	}
	// 定义快慢俩指针,当快指针到末尾的时候,慢指针肯定在链表中间位置
	slow,fast := head,head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow,fast = slow.Next,fast.Next.Next
	}
	// 把链表拆分成两段,所以设置中间位置即慢指针的next为nil
	n := slow.Next
	slow.Next = nil
	// 递归排序
	return merge(sortList(head),sortList(n))
}

func merge(node1 *ListNode,node2 *ListNode)*ListNode {
	// 设置一个空链表,
	node := &ListNode{Val:0}
	current := node
	// 挨个比较俩链表的值,把小的值放到新定义的链表里,排好序
	for node1 != nil && node2 != nil {
		if node1.Val <= node2.Val {
			current.Next,node1 = node1,node1.Next
		} else {
			current.Next,node2 = node2,node2.Next
		}
		current = current.Next
	}

	// 两链表可能有一个没走完,所以要把没走完的放到链表的后面
	// 注意,此处跟 数组不一样的是, 数组为什么要循环,因为数组可能一个数组全部走了(比如 12345与6789比较, 前面的全部走完,后面一个没走),另一个可能有多个没走..
	// 链表虽然也有这种可能,但是 node1和node2已经是有序的了,如果另外一个没有走完,直接把next指向node1或者node2就行,因为这是链表
	if node1 != nil {
		current.Next,node1 = node1,node1.Next
	}
	if node2 != nil {
		current.Next,node2 = node2,node2.Next
	}
	return node.Next
}
// @lc code=end

