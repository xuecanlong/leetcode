/*
 * @lc app=leetcode.cn id=1081 lang=golang
 *
 * [1081] 不同字符的最小子序列
 *
 * https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters/description/
 *
 * algorithms
 * Medium (41.56%)
 * Likes:    9
 * Dislikes: 0
 * Total Accepted:    830
 * Total Submissions: 2K
 * Testcase Example:  '"cdadabcc"'
 *
 * 返回字符串 text 中按字典序排列最小的子序列，该子序列包含 text 中所有不同字符一次。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入："cdadabcc"
 * 输出："adbc"
 * 
 * 
 * 示例 2：
 * 
 * 输入："abcd"
 * 输出："abcd"
 * 
 * 
 * 示例 3：
 * 
 * 输入："ecbacba"
 * 输出："eacb"
 * 
 * 
 * 示例 4：
 * 
 * 输入："leetcode"
 * 输出："letcod"
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= text.length <= 1000
 * text 由小写英文字母组成
 * 
 * 
 */

// @lc code=start

func smallestSubsequence(text string) string {
	/*
	   这题我学到了：
	    第一次遍历，先对字符计数
	    第二次遍历，按顺序依次减少剩余字符的数量，同时利用Stack弹出非最优解
	*/
	stack := make([]rune, len(text))
	endIndex := -1
	seen := map[rune]bool{}
	count := map[rune]int{}
	for _, c := range text {
		count[c]++
	}
	for _, c := range text {
		count[c]--
		if seen[c] {
			continue
		}
		for endIndex > -1 {
			v := stack[endIndex]
			if v > c && count[v] > 0 {
				endIndex--
				delete(seen, v)
			} else {
				break
			}
		}
		endIndex++
		seen[c] = true
		stack[endIndex] = c
	}
	return string(stack[:endIndex+1])

}
// @lc code=end

