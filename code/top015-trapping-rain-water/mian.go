package main

import (
	"fmt"
	"math"
)

//https://leetcode.cn/problems/trapping-rain-water/description/?envType=study-plan-v2&envId=top-interview-150
//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

//n == height.length
//1 <= n <= 2 * 104
//0 <= height[i] <= 105
func trap(height []int) int {
	ans, leftMax, RightMax := 0, 0, 0
	for idx, h := range height {
		leftMax = int(math.Max(float64(leftMax), float64(h)))
		RightMax = int(math.Max(float64(RightMax), float64(height[len(height)-idx-1])))
		ans += RightMax + leftMax - h
	}
	ans = ans - leftMax*len(height)
	return ans
}
func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
