package main

import "fmt"

// https://leetcode.cn/problems/majority-element/?envType=study-plan-v2&envId=top-interview-150
// 给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
// 输入：nums = [3,2,3]
// 输出：3
// 示例 2：
//
// 输入：nums = [2,2,1,1,1,2,2]
// 输出：2
//
// 提示：
// n == nums.length
// 1 <= n <= 5 * 104
// -10^9 <= nums[i] <= 10^9
func majorityElementV2(nums []int) int {
	cnts := make(map[int]int)
	for _, num := range nums {
		if _, ok := cnts[num]; ok {
			cnts[num] = cnts[num] + 1
		} else {
			cnts[num] = 1
		}
	}
	for i, cnt := range cnts {
		if cnt > len(nums)/2 {
			return i
		}
	}
	return -1
}
func majorityElement(nums []int) int {
	major := 0
	count := 0
	for _, num := range nums {
		if count == 0 {
			major = num
		}
		if num == major {
			count += 1
		} else {
			count -= 1
		}
	}
	return major
}

func main() {
	fmt.Println(majorityElementV2([]int{1, 2, 2, 2, 2}))
}
