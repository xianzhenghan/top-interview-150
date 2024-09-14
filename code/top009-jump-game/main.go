package main

import "fmt"

// 的最大长度。
//
// 判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。
//
// 示例 1：
//
// 输入：nums = [2,3,1,1,4]
// 输出：true
// 解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
// 示例 2：
//
// 输入：nums = [3,2,1,0,4]
// 输出：false
// 解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
//
// 提示：
//
// 1 <= nums.length <= 104
// 0 <= nums[i] <= 105
func canJump(nums []int) bool {
	dp := make([]int, len(nums))
	dp[0] = nums[0] //i之前最大跳跃次数
	if dp[0] <= 0 && len(nums) > 1 {
		return false
	}
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]-1, nums[i])
		if dp[i] <= 0 && i < len(nums)-1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canJump([]int{0, 1}))
}
