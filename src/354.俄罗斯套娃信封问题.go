/*
 * @lc app=leetcode.cn id=354 lang=golang
 *
 * [354] 俄罗斯套娃信封问题
 *
 * https://leetcode-cn.com/problems/russian-doll-envelopes/description/
 *
 * algorithms
 * Hard (32.06%)
 * Likes:    66
 * Dislikes: 0
 * Total Accepted:    4.3K
 * Total Submissions: 13.3K
 * Testcase Example:  '[[5,4],[6,4],[6,7],[2,3]]'
 *
 * 给定一些标记了宽度和高度的信封，宽度和高度以整数对形式 (w, h)
 * 出现。当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
 * 
 * 请计算最多能有多少个信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。
 * 
 * 说明:
 * 不允许旋转信封。
 * 
 * 示例:
 * 
 * 输入: envelopes = [[5,4],[6,4],[6,7],[2,3]]
 * 输出: 3 
 * 解释: 最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
 * 
 * 
 */

// @lc code=start
func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) <= 1 {
		return len(envelopes)
	}
	// 先按宽度排序
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	tail := make([]int, 0)
	for i := 0; i < len(envelopes); i++ {
		lo, li := 0, len(tail) - 1
		// 二分查找
		for lo <= li {
			mid := lo + (li - lo)/2
			if envelopes[i][1] > tail[mid] {
				lo = mid + 1
			} else {
				li = mid - 1
			}
		}
		if lo >= len(tail) {
			tail = append(tail, envelopes[i][1])
		} else {
			tail[lo] = envelopes[i][1]
		}
	}
	return len(tail)
}

func maxEnvelopes1(envelopes [][]int) int {
	if len(envelopes) <= 1 {
		return len(envelopes)
	}
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] < envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	max := 1
	dp := make([]int, len(envelopes))
	dp[0] = 1
	for i := 1; i < len(envelopes); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if envelopes[j][0] < envelopes[i][0] && envelopes[j][1] < envelopes[i][1] {
				if dp[i] < dp[j] + 1 {
					dp[i] = dp[j] + 1
				}
				if dp[i] > max {
					max = dp[i]
					break
				}
			}
		}
	}
	return max
}
// @lc code=end

