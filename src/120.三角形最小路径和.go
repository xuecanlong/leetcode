/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 *
 * https://leetcode-cn.com/problems/triangle/description/
 *
 * algorithms
 * Medium (62.19%)
 * Likes:    230
 * Dislikes: 0
 * Total Accepted:    25.8K
 * Total Submissions: 41.5K
 * Testcase Example:  '[[2],[3,4],[6,5,7],[4,1,8,3]]'
 *
 * 给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
 * 
 * 例如，给定三角形：
 * 
 * [
 * ⁠    [2],
 * ⁠   [3,4],
 * ⁠  [6,5,7],
 * ⁠ [4,1,8,3]
 * ]
 * 
 * 
 * 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
 * 
 * 说明：
 * 
 * 如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
 * 
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {

	lines := len(triangle)

	switch lines {
	case 0:
		return 0
	case 1:
		// 如果只有一行，则返回第一个元素
		return triangle[0][0]
	default:
	}

	// dp表格，用来保存 到达上一行各个节点的最小路径和
	// 该表中的元素一直重复读写
	dp := make([]int, len(triangle))

	// 初始化条件，第一行中的唯一元素
	dp[0] = triangle[0][0]

	var minTotal int

	for i := 1; i < lines; i++ {
		// 开始遍历所有的行

		currentLine := triangle[i]
		currentDP := make([]int, i+1)

		for j := 0; j < len(currentLine); j++ {
			// 处理当前行，j为第i行中的第j个元素

			// 如果是第0个元素，只有一条路径到
			if j == 0 {
				currentDP[j] = dp[j] + currentLine[j]
			} else if j == i {
				// 如果是最后一个元素
				currentDP[j] = dp[j-1] + currentLine[j]
			} else {
				// 否则取最小路径
				currentDP[j] = min(
					dp[j-1],
					dp[j],
				) + currentLine[j]
			}

			// 如果是最后一行，则开始记录最小值
			if i == lines-1 {
				if j == 0 {
					// 如果是第一个结果，则初始化返回结果
					minTotal = currentDP[j]
				} else {
					// 否则更新
					if minTotal > currentDP[j] {
						minTotal = currentDP[j]
					}
				}

			}
		}

        // 更新dp表
        dp = currentDP
	}

	return minTotal
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
// @lc code=end

