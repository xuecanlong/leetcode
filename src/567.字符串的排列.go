/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 *
 * https://leetcode-cn.com/problems/permutation-in-string/description/
 *
 * algorithms
 * Medium (32.35%)
 * Likes:    67
 * Dislikes: 0
 * Total Accepted:    9.1K
 * Total Submissions: 28.1K
 * Testcase Example:  '"ab"\n"eidbaooo"'
 *
 * 给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
 * 
 * 换句话说，第一个字符串的排列之一是第二个字符串的子串。
 * 
 * 示例1:
 * 
 * 
 * 输入: s1 = "ab" s2 = "eidbaooo"
 * 输出: True
 * 解释: s2 包含 s1 的排列之一 ("ba").
 * 
 * 
 * 
 * 
 * 示例2:
 * 
 * 
 * 输入: s1= "ab" s2 = "eidboaoo"
 * 输出: False
 * 
 * 
 * 
 * 
 * 注意：
 * 
 * 
 * 输入的字符串只包含小写字母
 * 两个字符串的长度都在 [1, 10,000] 之间
 * 
 * 
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	l1:=len(s1)
    l2:=len(s2)
    if l2<l1{
        return false
    }
    
    var c1  [26]int
    var c2  [26]int
    for k,v := range s1 {
        c1[v-'a']++
        c2[s2[k]-'a']++
    }
 
    for i:=0;i<l2-l1+1 ;i++  {
        if c1==c2 {
	return true
        }
        if i<l2-l1 {
			c2[s2[i]-'a']--
			c2[s2[i+l1]-'a']++
        }
    }
    return false
}
// @lc code=end

