/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 报数
 *
 * https://leetcode-cn.com/problems/count-and-say/description/
 *
 * algorithms
 * Easy (52.95%)
 * Likes:    326
 * Dislikes: 0
 * Total Accepted:    54.5K
 * Total Submissions: 102.9K
 * Testcase Example:  '1'
 *
 * 报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：
 * 
 * 1.     1
 * 2.     11
 * 3.     21
 * 4.     1211
 * 5.     111221
 * 
 * 
 * 1 被读作  "one 1"  ("一个一") , 即 11。
 * 11 被读作 "two 1s" ("两个一"）, 即 21。
 * 21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。
 * 
 * 给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。
 * 
 * 注意：整数顺序将表示为一个字符串。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: 1
 * 输出: "1"
 * 
 * 
 * 示例 2:
 * 
 * 输入: 4
 * 输出: "1211"
 * 
 * 
 */

// @lc code=start
// func countAndSay(n int) string {
//     var result string
//     result = "1" 
//     for i:=1; i<n; i++ {
//         var numbers []string
//         var counts []int
//         m:=0
//         for k:=0; k<len(result); k++ {
//             hasData := false
//             for ; m<len(numbers); m++ {
//                 if numbers[m]== string(result[k]) {
//                     counts[m]++
//                     hasData = true
//                     break
//                 }
//             }
//             if(!hasData) {
//                 numbers = append(numbers, string(result[k]))
//                 counts = append(counts, 1)
//             }
//         }
//         var buffer bytes.Buffer
//         for j:=0; j<len(numbers); j++ {
//             buffer.WriteString(fmt.Sprintf("%d", counts[j]))
//             buffer.WriteString(numbers[j])
//         }
//         result = buffer.String()
//     }
//     return result
// }

func countAndSay(n int) string {
	n--
	ret := "1"
	for i := 0; i < n; i++ {
		var str []rune
		l := len(ret)
		count := 49
		for j := 0; j < l; j++ {
			if j+1 == l {
				str = append(str, rune(count), rune(ret[j]))
				count = 49
			} else {
				if ret[j] == ret[j+1] {
					count++
				} else {
					str = append(str, rune(count), rune(ret[j]))
					count = 49
				}
			}
		}
		ret = string(str)
	}
	return ret
}
// @lc code=end

