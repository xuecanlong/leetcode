/*
 * @lc app=leetcode.cn id=331 lang=golang
 *
 * [331] 验证二叉树的前序序列化
 *
 * https://leetcode-cn.com/problems/verify-preorder-serialization-of-a-binary-tree/description/
 *
 * algorithms
 * Medium (42.08%)
 * Likes:    44
 * Dislikes: 0
 * Total Accepted:    2.6K
 * Total Submissions: 6.3K
 * Testcase Example:  '"9,3,4,#,#,1,#,#,2,#,6,#,#"'
 *
 * 序列化二叉树的一种方法是使用前序遍历。当我们遇到一个非空节点时，我们可以记录下这个节点的值。如果它是一个空节点，我们可以使用一个标记值记录，例如 #。
 * 
 * ⁠    _9_
 * ⁠   /   \
 * ⁠  3     2
 * ⁠ / \   / \
 * ⁠4   1  #  6
 * / \ / \   / \
 * # # # #   # #
 * 
 * 
 * 例如，上面的二叉树可以被序列化为字符串 "9,3,4,#,#,1,#,#,2,#,6,#,#"，其中 # 代表一个空节点。
 * 
 * 给定一串以逗号分隔的序列，验证它是否是正确的二叉树的前序序列化。编写一个在不重构树的条件下的可行算法。
 * 
 * 每个以逗号分隔的字符或为一个整数或为一个表示 null 指针的 '#' 。
 * 
 * 你可以认为输入格式总是有效的，例如它永远不会包含两个连续的逗号，比如 "1,,3" 。
 * 
 * 示例 1:
 * 
 * 输入: "9,3,4,#,#,1,#,#,2,#,6,#,#"
 * 输出: true
 * 
 * 示例 2:
 * 
 * 输入: "1,#"
 * 输出: false
 * 
 * 
 * 示例 3:
 * 
 * 输入: "9,#,#,1"
 * 输出: false
 * 
 */

// @lc code=start
// - 空节点当作树的叶子节点。
// - 树的特征：当遍历完整棵二叉树时，叶子节点的数目总是比非叶子节点的数目多1。
// - 从数组开头遍历，在遍历的过程中判断应保证非叶子节点的数目countNode不小于叶子节点的数目countNone。
// - 当遍历到最后一个位置时，需要保证最后一个元素为空，否则，该数组不可能是一个二叉树的前序序列化。
// - 另外，叶节点的数目不可能小于2。
func isValidSerialization(preorder string) bool {
    leaves := 0
    node := 0
    pres := strings.Split(preorder,",")
	for i, s := range pres {
		if s == "#" {
			leaves ++
		} else {
			node ++
		}
		if leaves > node + 1{
			return false

		}
		if leaves == node + 1 && i < len(pres)- 1 {
			return false
		}

	}
	if leaves != node + 1{
		return false
	} else {
		return true
	}
}
// @lc code=end

