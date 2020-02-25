/*
 * @lc app=leetcode.cn id=825 lang=golang
 *
 * [825] 适龄的朋友
 *
 * https://leetcode-cn.com/problems/friends-of-appropriate-ages/description/
 *
 * algorithms
 * Medium (30.89%)
 * Likes:    28
 * Dislikes: 0
 * Total Accepted:    1.4K
 * Total Submissions: 4.6K
 * Testcase Example:  '[16,16]'
 *
 * 人们会互相发送好友请求，现在给定一个包含有他们年龄的数组，ages[i] 表示第 i 个人的年龄。
 * 
 * 当满足以下条件时，A 不能给 B（A、B不为同一人）发送好友请求：
 * 
 * 
 * age[B] <= 0.5 * age[A] + 7
 * age[B] > age[A]
 * age[B] > 100 && age[A] < 100
 * 
 * 
 * 否则，A 可以给 B 发送好友请求。
 * 
 * 注意如果 A 向 B 发出了请求，不等于 B 也一定会向 A 发出请求。而且，人们不会给自己发送好友请求。 
 * 
 * 求总共会发出多少份好友请求?
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: [16,16]
 * 输出: 2
 * 解释: 二人可以互发好友申请。
 * 
 * 
 * 示例 2:
 * 
 * 输入: [16,17,18]
 * 输出: 2
 * 解释: 好友请求可产生于 17 -> 16, 18 -> 17.
 * 
 * 示例 3:
 * 
 * 输入: [20,30,100,110,120]
 * 输出: 3
 * 解释: 好友请求可产生于 110 -> 100, 120 -> 110, 120 -> 100.
 * 
 * 
 * 
 * 
 * 说明:
 * 
 * 
 * 1 <= ages.length <= 20000.
 * 1 <= ages[i] <= 120.
 * 
 * 
 */

// @lc code=start
func numFriendRequests(ages []int) int {
	ageCount := [121]int{}
	for i := 0; i < len(ages); i++ {
		ageCount[ages[i]]++ // 计算各年龄人数
	}
	for i := 1; i < 121; i++ {
		ageCount[i] += ageCount[i-1] // 优化：计算出年龄段的最大和最小值，直接 头 -（尾-1）得到区间内的总人数
	}
	ret := 0
	for i := 0; i < len(ages); i++ {
		low, high := int(0.5 * float64(ages[i])) + 7, ages[i] // 计算条件内的最大最小值
		if low >= high {
			continue
		}
		if ages[i] > low && ages[i] <= high { // 需要减去自身
			ret--
		}
		ret += ageCount[high] - ageCount[low] // 计算区间内的人数
	}
	return ret
}

// @lc code=end

