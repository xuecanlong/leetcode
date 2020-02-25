/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 *
 * https://leetcode-cn.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (46.28%)
 * Likes:    5717
 * Dislikes: 0
 * Total Accepted:    449.4K
 * Total Submissions: 971K
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
 * 
 * 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
 * 
 * 示例:
 * 
 * 给定 nums = [2, 7, 11, 15], target = 9
 * 
 * 因为 nums[0] + nums[1] = 2 + 7 = 9
 * 所以返回 [0, 1]
 * 
 * 
 */
func twoSum(nums []int, target int) []int {
    lookup := make(map[int]int)
    for i, v := range nums {
        j, ok := lookup[v]
        lookup[target - v] = i
        if ok {
            return []int{j, i}
        }
    }
    return []int{}
}


