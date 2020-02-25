/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 *
 * https://leetcode-cn.com/problems/single-number/description/
 *
 * algorithms
 * Easy (63.76%)
 * Likes:    889
 * Dislikes: 0
 * Total Accepted:    114K
 * Total Submissions: 178.5K
 * Testcase Example:  '[2,2,1]'
 *
 * 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
 * 
 * 说明：
 * 
 * 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
 * 
 * 示例 1:
 * 
 * 输入: [2,2,1]
 * 输出: 1
 * 
 * 
 * 示例 2:
 * 
 * 输入: [4,1,2,1,2]
 * 输出: 4
 * 
 */

// @lc code=start
// func singleNumber(nums []int) int {
// 	m := make(map[int]int)

// 	for _, data := range nums {
// 		if _, ok := m[data]; ok {
// 			delete(m, data)
// 		} else {
// 			m[data] = 1
// 		}
// 	}
// 	for k := range m {
// 		return k
// 	}
// 	return 0
// }

// 如果我们对 0 和二进制位做 XOR (异或)运算，得到的仍然是这个二进制位
// aa⊕0=a
// 如果我们对相同的二进制位做 XOR 运算，返回的结果是 0
// 0a⊕a=0
// XOR 满足交换律和结合律
// ba⊕b⊕a=(a⊕a)⊕b=0⊕b=b
// 所以我们只需要将所有的数进行 XOR 操作，得到那个唯一的数字。
func singleNumber(nums []int) int {
	res := 0
	for _, data := range nums {
		res ^= data
	}
	return res
}
// @lc code=end

