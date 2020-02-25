/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (38.83%)
 * Likes:    911
 * Dislikes: 0
 * Total Accepted:    93.7K
 * Total Submissions: 241.3K
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 * 
 * 有效字符串需满足：
 * 
 * 
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 * 
 * 
 * 注意空字符串可被认为是有效字符串。
 * 
 * 示例 1:
 * 
 * 输入: "()"
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入: "()[]{}"
 * 输出: true
 * 
 * 
 * 示例 3:
 * 
 * 输入: "(]"
 * 输出: false
 * 
 * 
 * 示例 4:
 * 
 * 输入: "([)]"
 * 输出: false
 * 
 * 
 * 示例 5:
 * 
 * 输入: "{[]}"
 * 输出: true
 * 
 */
 func isValid(s string) bool {
    if s == "" {
        return true
    }
    
    tr := map[rune]rune{
        '}': '{',
        ')': '(',
        ']': '[',
    }
    
    stack := make([]rune, 0, 1)
    
    for _, ch := range s {
        switch ch {
        case 40,123,91:
            stack = append(stack, ch)
        case 41,125,93:    
            if len(stack) == 0 || stack[len(stack)-1] != tr[ch] {
                return false
            } else {
                stack = stack[:len(stack)-1]
            }
        }
    }
    
    return len(stack) == 0
}
