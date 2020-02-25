/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 *
 * https://leetcode-cn.com/problems/intersection-of-two-arrays/description/
 *
 * algorithms
 * Easy (66.44%)
 * Likes:    130
 * Dislikes: 0
 * Total Accepted:    37.5K
 * Total Submissions: 56.2K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * 给定两个数组，编写一个函数来计算它们的交集。
 * 
 * 示例 1:
 * 
 * 输入: nums1 = [1,2,2,1], nums2 = [2,2]
 * 输出: [2]
 * 
 * 
 * 示例 2:
 * 
 * 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * 输出: [9,4]
 * 
 * 说明:
 * 
 * 
 * 输出结果中的每个元素一定是唯一的。
 * 我们可以不考虑输出结果的顺序。
 * 
 * 
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
    var count = map[int]bool{}
    var result = []int{}
    for _,num := range nums1{
        count[num]=true
    }
    for _, num := range nums2{
        if(count[num]==true){
            result=append(result,num)
            count[num]=false
        }
    } 
    return result    
}
// @lc code=end

