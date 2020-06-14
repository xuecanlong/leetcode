/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 *
 * https://leetcode-cn.com/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (65.00%)
 * Likes:    609
 * Dislikes: 0
 * Total Accepted:    108.2K
 * Total Submissions: 166.4K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 反转一个单链表。
 * 
 * 示例:
 * 
 * 输入: 1->2->3->4->5->NULL
 * 输出: 5->4->3->2->1->NULL
 * 
 * 进阶:
 * 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
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
 func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}

// 链表反转口诀：斩断后路,不忘前事,才能重获新生,非递归
func reverseList(head *ListNode) *ListNode {
    var pre *ListNode = nil
    cur := head
    for nil != cur {
        // 1.（保存一下前进方向）保存下一跳
        temp := cur.Next
        // 2.斩断过去,不忘前事
        cur.Next = pre
        // 3.前驱指针的使命在上面已经完成，这里需要更新前驱指针
        pre = cur
        // 当前指针的使命已经完成，需要继续前进了
        cur = temp
    }
    return pre
}

// 递归
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    newHead := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return newHead
}
// @lc code=end

