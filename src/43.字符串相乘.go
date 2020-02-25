/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 *
 * https://leetcode-cn.com/problems/multiply-strings/description/
 *
 * algorithms
 * Medium (40.71%)
 * Likes:    218
 * Dislikes: 0
 * Total Accepted:    30K
 * Total Submissions: 73.6K
 * Testcase Example:  '"2"\n"3"'
 *
 * 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
 * 
 * 示例 1:
 * 
 * 输入: num1 = "2", num2 = "3"
 * 输出: "6"
 * 
 * 示例 2:
 * 
 * 输入: num1 = "123", num2 = "456"
 * 输出: "56088"
 * 
 * 说明：
 * 
 * 
 * num1 和 num2 的长度小于110。
 * num1 和 num2 只包含数字 0-9。
 * num1 和 num2 均不以零开头，除非是数字 0 本身。
 * 不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
 * 
 * 
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	num1, num2 = "0"+num1, "0"+num2 // sentinel to avoid special handling for the last carry over
	res := make([]byte, len(num1)+len(num2))
	for i, _ := range res { // initialize with byte character '0'
		res[i] = '0'
	}

	for i := 0; i < len(num2); i++ {
		pos2 := len(num2) - 1 - i
		ch2 := num2[pos2]
		co := 0
		for j := 0; j < len(num1); j++ {
			pos1 := len(num1) - 1 - j
			ch1 := num1[pos1]
			resPos := len(res) - 1 - (i + j)

			mul := int(ch1-'0')*int(ch2-'0') + int(res[resPos]-'0') + co
			res[resPos] = byte(mul%10 + '0')
			co = mul / 10
		}
	}

	var b bytes.Buffer
	firstZero := true
	for i := 0; i < len(res); i++ {
		c := res[i]
		if firstZero && c == '0' {
			continue
		}
		firstZero = false
		b.WriteByte(c)
	}

	resStr := b.String()
	if resStr == "" {
		return "0"
	}
	return resStr
}
// @lc code=end

