/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 *
 * https://leetcode-cn.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (41.21%)
 * Likes:    121
 * Dislikes: 0
 * Total Accepted:    59.1K
 * Total Submissions: 143K
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
 * 
 * 说明：本题中，我们将空字符串定义为有效的回文串。
 * 
 * 示例 1:
 * 
 * 输入: "A man, a plan, a canal: Panama"
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入: "race a car"
 * 输出: false
 * 
 * 
 */

// @lc code=start
func Reverse(s string) string {
    n := len(s)
    runes := make([]rune, n)
    for _, rune := range s {
      n--
      runes[n] = rune
    }
    return string(runes[n:])
}


func isPalindrome(s string) bool {
    isalphanum := func(r rune) rune {
        if unicode.IsLetter(r) || unicode.IsNumber(r) {
            return r
        }
        return -1
    }
    s2 := strings.ToLower(strings.Map(isalphanum, s))
    return s2 == Reverse(s2)
}
// @lc code=end

