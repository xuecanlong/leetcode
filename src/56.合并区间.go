/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 *
 * https://leetcode-cn.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (38.90%)
 * Likes:    198
 * Dislikes: 0
 * Total Accepted:    32.7K
 * Total Submissions: 84K
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * 给出一个区间的集合，请合并所有重叠的区间。
 * 
 * 示例 1:
 * 
 * 输入: [[1,3],[2,6],[8,10],[15,18]]
 * 输出: [[1,6],[8,10],[15,18]]
 * 解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 * 
 * 
 * 示例 2:
 * 
 * 输入: [[1,4],[4,5]]
 * 输出: [[1,5]]
 * 解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
 * 
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	sort.Sort(ByStart(intervals))
	interLen := len(intervals)

	for i := 0; i < interLen-1; {
		cur, next := intervals[i], intervals[i+1]
		if next[0] <= cur[1] {
			if next[1] > cur[1] {
				cur[1] = next[1]
			}
			intervals = append(intervals[:i+1], intervals[i+2:]...)
			interLen--
			continue
		}
		i++
	}
	return intervals
}
type Interval struct {
	Start int
	End int
}

type ByStart [][]int

func (bs ByStart) Len() int {
	return len(bs)
}
func (bs ByStart) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}
func (bs ByStart) Less(i, j int) bool {
	return bs[i][0] < bs[j][0]
}
// @lc code=end

