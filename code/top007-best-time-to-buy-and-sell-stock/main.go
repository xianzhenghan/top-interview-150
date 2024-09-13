package main

import (
	"fmt"
)

// 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
//
// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
//
// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
//
// 示例 1：
//
// 输入：[7,1,5,3,6,4]
// 输出：5
// 解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//
//	注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
//
// 示例 2：
//
// 输入：prices = [7,6,4,3,1]
// 输出：0
// 解释：在这种情况下, 没有交易完成, 所以最大利润为 0。
//
// 提示：
//
// 1 <= prices.length <= 105
// 0 <= prices[i] <= 104
func maxProfitv2(prices []int) int {
	maxProfit := 0
	minPrice := prices[0]
	for _, price := range prices {
		minPrice = min(minPrice, price)
		maxProfit = max(maxProfit, price-minPrice)
	}
	return maxProfit
}

func maxProfit(prices []int) int {
	maxProfit := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > maxProfit {
				maxProfit = prices[j] - prices[i]
			}
		}
	}
	return maxProfit
}

//变种题目
// best-time-to-buy-and-sell-stock-ii/
// 给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
//
// 在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
//
// 返回 你能获得的 最大 利润 。
//例 1：
//
//输入：prices = [7,1,5,3,6,4]
//输出：7
//解释：在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4。
//随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6 - 3 = 3。
//最大总利润为 4 + 3 = 7 。
//1 <= prices.length <= 3 * 104
//0 <= prices[i] <= 104

const FREE, HOLD = 0, 1

func maxProfitii(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][FREE] = 0
	dp[0][HOLD] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][FREE] = max(dp[i-1][FREE], dp[i-1][HOLD]+prices[i])
		dp[i][HOLD] = max(dp[i-1][HOLD], dp[i-1][FREE]-prices[i])
	}
	return dp[n-1][FREE]
}
func maxProfitiiV2(prices []int) int {
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}
func main() {
	fmt.Println(maxProfitiiV2([]int{7, 1, 5, 3, 6, 4}))
}
