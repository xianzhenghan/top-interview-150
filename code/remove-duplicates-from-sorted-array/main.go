package main

import "fmt"

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/submissions/563657147/?envType=study-plan-v2&envId=top-interview-150
// 1 <= nums.length <= 3 * 104
// -104 <= nums[i] <= 104
func removeDuplicates(nums []int) int {
	dp := make([]int, 1)
	dp[0] = nums[0]
	for i, j := 1, 1; i < len(nums); i++ {
		if dp[j-1] != nums[i] {
			dp = append(dp, nums[i])
			j++
		}
	}
	copy(nums[:len(dp)], dp)
	return len(dp)
}

func removeDuplicatesV2(nums []int) int {
	slow, fast := 1, 1
	for ; fast < len(nums); fast++ {
		if nums[fast] != nums[slow-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func main() {
	fmt.Println(removeDuplicatesV2([]int{1}))
}
