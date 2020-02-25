/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 *
 * https://leetcode-cn.com/problems/max-area-of-island/description/
 *
 * algorithms
 * Medium (56.79%)
 * Likes:    147
 * Dislikes: 0
 * Total Accepted:    11.1K
 * Total Submissions: 19.5K
 * Testcase Example:  '[[1,1,0,0,0],[1,1,0,0,0],[0,0,0,1,1],[0,0,0,1,1]]'
 *
 * 给定一个包含了一些 0 和 1的非空二维数组 grid , 一个 岛屿 是由四个方向 (水平或垂直) 的 1 (代表土地)
 * 构成的组合。你可以假设二维矩阵的四个边缘都被水包围着。
 * 
 * 找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)
 * 
 * 示例 1:
 * 
 * 
 * [[0,0,1,0,0,0,0,1,0,0,0,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,1,0,0,0],
 * ⁠[0,1,1,0,1,0,0,0,0,0,0,0,0],
 * ⁠[0,1,0,0,1,1,0,0,1,0,1,0,0],
 * ⁠[0,1,0,0,1,1,0,0,1,1,1,0,0],
 * ⁠[0,0,0,0,0,0,0,0,0,0,1,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,1,0,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,0,0,0,0]]
 * 
 * 
 * 对于上面这个给定矩阵应返回 6。注意答案不应该是11，因为岛屿只能包含水平或垂直的四个方向的‘1’。
 * 
 * 示例 2:
 * 
 * 
 * [[0,0,0,0,0,0,0,0]]
 * 
 * 对于上面这个给定的矩阵, 返回 0。
 * 
 * 注意: 给定的矩阵grid 的长度和宽度都不超过 50。
 * 
 */

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
    maxArea := 0
    for y, maxY := 0, len(grid); y < maxY; y++ {
        for x, maxX := 0, len(grid[y]); x < maxX; x++ {
            if grid[y][x] > 0 {
                if sum := sumIslandArea(grid, y, x); sum > maxArea {
                    maxArea = sum
                }
            }
        }
    }
    return maxArea
}

func sumIslandArea(grid [][]int, y int, x int) int {
    sum := 1
    grid[y][x] = 0

    maxY, maxX := len(grid), len(grid[y])
    if y > 0 && grid[y-1][x] > 0 {
        sum += sumIslandArea(grid, y-1, x)
    }
    if x < maxX-1 && grid[y][x+1] > 0 {
        sum += sumIslandArea(grid, y, x+1)
    }
    if y < maxY-1 && grid[y+1][x] > 0 {
        sum += sumIslandArea(grid, y+1, x)
    }
    if x > 0 && grid[y][x-1] > 0 {
        sum += sumIslandArea(grid, y, x-1)
    }
    return sum
}
// @lc code=end

