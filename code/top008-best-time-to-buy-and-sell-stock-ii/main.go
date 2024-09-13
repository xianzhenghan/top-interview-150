package main

import "fmt"

func maxProfit(prices []int) int {
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}

// 使用动态规划
func maxProfit_dp(prices []int) int {
	return 0
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
