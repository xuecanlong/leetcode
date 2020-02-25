/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 第k个排列
 *
 * https://leetcode-cn.com/problems/permutation-sequence/description/
 *
 * algorithms
 * Medium (47.17%)
 * Likes:    132
 * Dislikes: 0
 * Total Accepted:    15.3K
 * Total Submissions: 32.5K
 * Testcase Example:  '3\n3'
 *
 * 给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。
 * 
 * 按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
 * 
 * 
 * "123"
 * "132"
 * "213"
 * "231"
 * "312"
 * "321"
 * 
 * 
 * 给定 n 和 k，返回第 k 个排列。
 * 
 * 说明：
 * 
 * 
 * 给定 n 的范围是 [1, 9]。
 * 给定 k 的范围是[1,  n!]。
 * 
 * 
 * 示例 1:
 * 
 * 输入: n = 3, k = 3
 * 输出: "213"
 * 
 * 
 * 示例 2:
 * 
 * 输入: n = 4, k = 9
 * 输出: "2314"
 * 
 * 
 */

// @lc code=start
func getPermutation(n int, k int) string {
	var nums []int
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	return getSolution(nums, "", n, k-1)
}

func getSolution(nums []int, str string, n int, k int) string {
	if n == 1 {
		return str + strconv.Itoa(nums[0])
	}
	//当有n位集合时，可以知道第一位显然有n个可能，而每一个第一位确定后，引申出来的可能性有 (n-1)! 个
	//所以反过来呢，第k个排列的第一位，就是 k/(n-1)! 余数记为rest
	//于是第二位的答案也呼之欲出： rest/(n-2)!
	part := factorial(n - 1)
	startNumberIndex := k / part
	rest := k % part
	startNumber := nums[startNumberIndex]
	nums = append(nums[:startNumberIndex], nums[startNumberIndex+1:]...)

	return getSolution(nums, str+strconv.Itoa(startNumber), n-1, rest)
}

//计算 n!
func factorial(num int) int {
	if num <= 1 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}
// @lc code=end

