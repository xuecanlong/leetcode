/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 *
 * https://leetcode-cn.com/problems/happy-number/description/
 *
 * algorithms
 * Easy (55.40%)
 * Likes:    176
 * Dislikes: 0
 * Total Accepted:    32.4K
 * Total Submissions: 58.3K
 * Testcase Example:  '19'
 *
 * 编写一个算法来判断一个数是不是“快乐数”。
 * 
 * 一个“快乐数”定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是无限循环但始终变不到
 * 1。如果可以变为 1，那么这个数就是快乐数。
 * 
 * 示例: 
 * 
 * 输入: 19
 * 输出: true
 * 解释: 
 * 1^2 + 9^2 = 82
 * 8^2 + 2^2 = 68
 * 6^2 + 8^2 = 100
 * 1^2 + 0^2 + 0^2 = 1
 * 
 * 
 */

// @lc code=start
func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	if n == 0 || n == 4 {
		return false
	}
	s := strconv.Itoa(n)
	res := 0
	for _, data := range s {
		num := int(data - '1' + 1)
		res += num * num
	}
	return isHappy(res)
}

// func isHappy(n int) bool {
// 	hash := make(map[int]bool)      // 开一个 map 判断是否循环
// 	for n!=1 {
// 		if _,ok := hash[n]; ok {    // 出现了循环，证明不是快乐数
// 			return false
// 		}
// 		hash[n] = true
// 		next := 0
// 		for n!=0 {                  // 计算下一个数字
// 			next += (n%10) * (n%10)
// 			n /= 10
// 		}
// 		n = next
// 	}
// 	return true
// }
// @lc code=end

