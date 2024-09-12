package main

import "fmt"

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/?envType=study-plan-v2&envId=top-interview-150

// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。
//
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
//1 <= nums.length <= 3 * 104
//-104 <= nums[i] <= 104
//nums 已按升序排列

// 输入：nums = [1,1,1,2,2,3]
// 输出：5, nums = [1,1,2,2,3]
// 解释：函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3。 不需要考虑数组中超出新长度后面的元素。
func removeDuplicates(nums []int) int {
	slow, fast, cnt := 1, 1, 1
	for ; fast < len(nums); fast++ {
		if nums[fast] != nums[slow-1] {
			nums[slow] = nums[fast]
			slow++
			cnt = 1
		} else {
			cnt++
			if cnt <= 2 {
				nums[slow] = nums[fast]
				slow++
			}
		}
		fmt.Println(slow, fast, cnt, nums)
	}
	fmt.Println(slow, fast, cnt, nums)
	return slow
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 1, 1, 1, 2, 3, 3}))
}
