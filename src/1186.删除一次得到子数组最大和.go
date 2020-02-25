/*
 * @lc app=leetcode.cn id=1186 lang=golang
 *
 * [1186] 删除一次得到子数组最大和
 *
 * https://leetcode-cn.com/problems/maximum-subarray-sum-with-one-deletion/description/
 *
 * algorithms
 * Medium (23.47%)
 * Likes:    11
 * Dislikes: 0
 * Total Accepted:    1.1K
 * Total Submissions: 4.6K
 * Testcase Example:  '[1,-2,0,3]'
 *
 * 给你一个整数数组，返回它的某个 非空 子数组（连续元素）在执行一次可选的删除操作后，所能得到的最大元素总和。
 * 
 * 
 * 换句话说，你可以从原数组中选出一个子数组，并可以决定要不要从中删除一个元素（只能删一次哦），（删除后）子数组中至少应当有一个元素，然后该子数组（剩下）的元素总和是所有子数组之中最大的。
 * 
 * 注意，删除一个元素后，子数组 不能为空。
 * 
 * 请看示例：
 * 
 * 示例 1：
 * 
 * 输入：arr = [1,-2,0,3]
 * 输出：4
 * 解释：我们可以选出 [1, -2, 0, 3]，然后删掉 -2，这样得到 [1, 0, 3]，和最大。
 * 
 * 示例 2：
 * 
 * 输入：arr = [1,-2,-2,3]
 * 输出：3
 * 解释：我们直接选出 [3]，这就是最大和。
 * 
 * 
 * 示例 3：
 * 
 * 输入：arr = [-1,-1,-1,-1]
 * 输出：-1
 * 解释：最后得到的子数组不能为空，所以我们不能选择 [-1] 并从中删去 -1 来得到 0。
 * ⁠    我们应该直接选择 [-1]，或者选择 [-1, -1] 再从中删去一个 -1。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= arr.length <= 10^5
 * -10^4 <= arr[i] <= 10^4
 * 
 * 
 */

// @lc code=start
// * 1.最大连续子数组和（noDel）的求法是:遍历到i时，以i为结尾元素的子数组，如果以i-1为结尾元素的子数组和大于0，那么以i为结尾元素的子数组可以把以i-1为结尾元素的子数组加进来作为数组开头，否则以i为结尾元素的连续子数组只有i一个元素；
// * 2.求删除一个元素的情况下的最大连续子数组和（del）的求法是:遍历到i时，要么删除arr[i],当前值也就变成了noDel[i-1],要么删除的是前面i-1各元素中的某一个，当前值就变成了del[i-1]+arr[i];
// * 3.觉得没必要保留noDel和del数组，都用单个变量就行了，反正只用i-1的记录，不会用i-2的记录
func maximumSum(arr []int) int {
	//删的最大连续子数组和，只有一个元素时不能删唯一的元素，所以把初始值置为最小值
	del := -math.MaxInt16
	//不删的最大连续子数组和
	noDel := arr[0]
	//初始化ans
	ans := del
	if noDel > del {
		ans = noDel
	}
	for i := 1; i < len(arr); i++ {
		if del+arr[i] > noDel {
			del = del+arr[i]
		} else {
			del = noDel
		}
		if noDel+arr[i] > arr[i] {
			noDel = noDel+arr[i]
		} else {
			noDel = arr[i]
		}
		if noDel > ans && noDel > del {
			ans = noDel
		} else if del > ans && del > noDel {
			ans = del
		}
	}
	return ans
}
// @lc code=end

