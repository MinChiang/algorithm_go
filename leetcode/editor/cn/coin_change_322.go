//给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
//
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
//
// 你可以认为每种硬币的数量是无限的。
//
//
//
// 示例 1：
//
//
//输入：coins = [1, 2, 5], amount = 11
//输出：3
//解释：11 = 5 + 5 + 1
//
// 示例 2：
//
//
//输入：coins = [2], amount = 3
//输出：-1
//
// 示例 3：
//
//
//输入：coins = [1], amount = 0
//输出：0
//
//
//
//
// 提示：
//
//
// 1 <= coins.length <= 12
// 1 <= coins[i] <= 2³¹ - 1
// 0 <= amount <= 10⁴
//
//
// Related Topics 广度优先搜索 数组 动态规划 👍 2203 👎 0

package main

func main() {
	coins := []int{1, 2, 5}
	print(coinChange(coins, 11))
}

//leetcode submit region begin(Prohibit modification and deletion)

//设dp(n)为：在由coins组成的总金额n，所使用的【最少】的硬币数量
//dp(0) = 0
//dp(n) = min{ dp(n - coins(i)) + 1 }
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		count := -1
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				// 说明可取
				if count == -1 {
					count = dp[i-coin] + 1
				} else if dp[i-coin]+1 < count {
					count = dp[i-coin] + 1
				}
			}
		}
		dp[i] = count
	}
	return dp[amount]
}

//leetcode submit region end(Prohibit modification and deletion)
