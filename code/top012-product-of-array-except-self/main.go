package main

import "fmt"

// https://leetcode.cn/problems/product-of-array-except-self/?envType=study-plan-v2&envId=top-interview-150
// 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
//
// 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
//
// 请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
//
// 示例 1:
//
// 输入: nums = [1,2,3,4]
// 输出: [24,12,8,6]
// 示例 2:
//
// 输入: nums = [-1,1,0,-3,3]
// 输出: [0,0,9,0,0]
//
// 提示：
//
// 2 <= nums.length <= 105
// -30 <= nums[i] <= 30
// 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内
// prefix and suffix
// 进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组 不被视为 额外空间。）
func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums)) //前i数据乘积(不包含i)
	prefix[0] = 1
	prefix[1] = nums[0]
	suffix := make([]int, len(nums))
	suffix[len(nums)-1] = 1                 //后i+1,...n 数据乘积(不包含i)
	suffix[len(nums)-2] = nums[len(nums)-1] //后i+1,...n 数据乘积(不包含i)
	for i := 2; i < len(nums); i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}
	for i := len(nums) - 3; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = prefix[i] * suffix[i]
	}
	return dp
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
