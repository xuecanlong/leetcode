/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 *
 * https://leetcode-cn.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (37.06%)
 * Likes:    240
 * Dislikes: 0
 * Total Accepted:    65.7K
 * Total Submissions: 177.1K
 * Testcase Example:  '4'
 *
 * 实现 int sqrt(int x) 函数。
 * 
 * 计算并返回 x 的平方根，其中 x 是非负整数。
 * 
 * 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
 * 
 * 示例 1:
 * 
 * 输入: 4
 * 输出: 2
 * 
 * 
 * 示例 2:
 * 
 * 输入: 8
 * 输出: 2
 * 说明: 8 的平方根是 2.82842..., 
 * 由于返回类型是整数，小数部分将被舍去。
 * 
 * 
 */

// @lc code=start
// func mySqrt(x int) int {
// 	//牛顿迭代法
// 	r := x
// 	for r*r > x {
// 		r = (r + x/r) / 2
// 	}
// 	return r
// }

func mySqrt(x int) int {
    // 二分查找，定位平方根
    i,j := 1,x
    for i<=j {
        mid := (i+j)/2
        if mid*mid > x {        // 去左区间找
            j = mid - 1
        } else if mid*mid < x { // 去右区间找
            i = mid + 1
        } else {                // 找到了
            return mid
        }
    }
    return i-1
}
// @lc code=end

