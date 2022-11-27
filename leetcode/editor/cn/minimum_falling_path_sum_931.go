//给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。
//
// 下降路径 可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第
//一个元素）。具体来说，位置 (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者 (row + 1
//, col + 1) 。
//
//
//
// 示例 1：
//
//
//
//
//输入：matrix = [[2,1,3],[6,5,4],[7,8,9]]
//输出：13
//解释：如图所示，为和最小的两条下降路径
//
//
// 示例 2：
//
//
//
//
//输入：matrix = [[-19,57],[-40,-5]]
//输出：-59
//解释：如图所示，为和最小的下降路径
//
//
//
//
// 提示：
//
//
// n == matrix.length == matrix[i].length
// 1 <= n <= 100
// -100 <= matrix[i][j] <= 100
//
//
// Related Topics 数组 动态规划 矩阵 👍 201 👎 0

package main

func main() {
	test := [][]int{
		{100, -42, -46, -41},
		{31, 97, 10, -10},
		{-58, -51, 82, 89},
		{51, 81, 69, -51},
	}
	print(minFallingPathSum(test))
}

//leetcode submit region begin(Prohibit modification and deletion)
// dp(i, j) = min{ dp(i - 1, j), dp(i - 1, j + 1)} + matrix(i, j)   i > 0, j = 0
// dp(i, j) = min{ dp(i - 1, j), dp(i - 1, j - 1)} + matrix(i, j)   i > 0, j = len - 1
// dp(i, j) = min{ dp(i - 1, j), dp(i - 1, j - 1), dp(i - 1, j + 1) } + matrix(i, j)   i > 0, 0 < j < len - 1
func minFallingPathSum(matrix [][]int) int {
	dp := make([]int, len(matrix))
	tmp := make([]int, len(matrix))
	copy(dp, matrix[0])
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if j == 0 {
				tmp[j] = getMin(dp[j], dp[j+1]) + matrix[i][j]
			} else if j == len(matrix)-1 {
				tmp[j] = getMin(dp[j], dp[j-1]) + matrix[i][j]
			} else {
				tmp[j] = getMin(dp[j], dp[j-1], dp[j+1]) + matrix[i][j]
			}
		}
		copy(dp, tmp)
	}
	return getMin(dp...)
}

func getMin(values ...int) int {
	min := values[0]
	for i := 1; i < len(values); i++ {
		if min > values[i] {
			min = values[i]
		}
	}
	return min
}

//leetcode submit region end(Prohibit modification and deletion)
