/*
 * @lc app=leetcode.cn id=1046 lang=golang
 *
 * [1046] 最后一块石头的重量
 *
 * https://leetcode-cn.com/problems/last-stone-weight/description/
 *
 * algorithms
 * Easy (57.53%)
 * Likes:    16
 * Dislikes: 0
 * Total Accepted:    6.1K
 * Total Submissions: 10.6K
 * Testcase Example:  '[2,7,4,1,8,1]'
 *
 * 有一堆石头，每块石头的重量都是正整数。
 * 
 * 每一回合，从中选出两块最重的石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：
 * 
 * 
 * 如果 x == y，那么两块石头都会被完全粉碎；
 * 如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
 * 
 * 
 * 最后，最多只会剩下一块石头。返回此石头的重量。如果没有石头剩下，就返回 0。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= stones.length <= 30
 * 1 <= stones[i] <= 1000
 * 
 * 
 */

// @lc code=start
func lastStoneWeight(stones []int) int {
	max1, max2 := 0, 0
	max1Index, max2Index := 0, 0
	for i, stone := range stones {
		if stone > max1 {
			max2 = max1
			max1 = stone
			max2Index = max1Index
			max1Index = i
		} else if stone > max2 {
			max2 = stone
			max2Index = i
		}
	}
	if max1 == 0 {
		return 0
	}
	if max2 == 0 {
		return max1
	}
	stones[max1Index], stones[max2Index] = max1-max2, 0
	return lastStoneWeight(stones)
}
// @lc code=end

