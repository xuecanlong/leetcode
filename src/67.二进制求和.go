/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 *
 * https://leetcode-cn.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (51.09%)
 * Likes:    252
 * Dislikes: 0
 * Total Accepted:    45.3K
 * Total Submissions: 88.6K
 * Testcase Example:  '"11"\n"1"'
 *
 * 给定两个二进制字符串，返回他们的和（用二进制表示）。
 * 
 * 输入为非空字符串且只包含数字 1 和 0。
 * 
 * 示例 1:
 * 
 * 输入: a = "11", b = "1"
 * 输出: "100"
 * 
 * 示例 2:
 * 
 * 输入: a = "1010", b = "1011"
 * 输出: "10101"
 * 
 */

// @lc code=start
func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	res := ""
	flag := 0
	for i >= 0 || j >= 0 {
		bitA, bitB := 0, 0
		if i >= 0 {
			bitA = int(a[i] - '0')
		}
		if j >= 0 {
			bitB = int(b[j] - '0')
		}
		current := bitA + bitB + flag
		flag = 0
		if current > 1 {
			flag = 1
			current = current -2
		}
		res = strconv.Itoa(current) + res
		i--
		j--
	}
	if flag == 1 {
		 res = "1" + res
	}
	return res
}
// @lc code=end

