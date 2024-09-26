package main

import (
	"fmt"
	"math"
	"runtime"
)

// 中等
// 相关标签
// 相关企业
// 给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。
//
// 每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:
//
// 0 <= j <= nums[i]
// i + j < n
// 返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。
//
// 示例 1:
//
// 输入: nums = [2,3,1,1,4]
// 输出: 2
// 解释: 跳到最后一个位置的最小跳跃数是 2。
//
//	从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
//
// 示例 2:
//
// 输入: nums = [2,3,0,1,4]
// 输出: 2
//
// 提示:
//
// 1 <= nums.length <= 104
// 0 <= nums[i] <= 1000
// 题目保证可以到达 nums[n-1]

func jumpV2(nums []int) int {
	// 初始化变量
	// 维护能够跳到的最大索引
	maxJump := 0
	// 维护某一跳所能达到的最大索引
	end := 0
	// 跳跃次数
	res := 0

	// 遍历数组，从索引0开始跳跃
	n := len(nums)
	for i := 0; i < n-1; i++ {
		// 每次跳跃都可以到达一个最大索引，即每次跳跃可以覆盖起跳点到最大索引的这一部分索引
		// 根据本次跳跃覆盖的这一部分索引，逐个计算并比较出下一次跳跃所能达到的最大索引
		maxJump = max(maxJump, i+nums[i])
		// 当索引到达本次跳跃的最大索引时，说明该进行下一次跳跃
		if i == end {
			// 更新下一次跳跃所能得到的最大索引
			end = maxJump
			// 并将跳跃次数加1，表示开始下一次跳跃
			res++
		}
		// 如果下一次跳跃的最大索引，超过了最后一个元素的索引，则可以提前退出
		if end >= n-1 {
			break
		}
	}

	// 返回跳跃次数
	return res
}

func jump(nums []int) int {
	if len(nums) == 1 {
		if nums[0] <= 1 {
			return 0
		} else {
			return 1
		}
	}
	minV := math.MaxInt
	dfs(0, 0, &minV, nums)
	return minV
}
func dfs(idx, count int, minV *int, nums []int) {
	for setp := 1; setp <= nums[idx]; setp++ {
		if idx+setp < len(nums)-1 {
			dfs(idx+setp, count+1, minV, nums)
		} else if idx+setp >= len(nums)-1 {
			fmt.Println(idx+setp, count)
			if count+1 < *minV {
				*minV = count + 1
			}
		}
	}
}

func main() {
	//模拟用户需求业务的数量
	task_cnt := math.MaxInt64

	for i := 0; i < task_cnt; i++ {
		go func(i int) {
			//... do some busi...

			fmt.Println("go func ", i, " goroutine count = ", runtime.NumGoroutine())
		}(i)
	}
}
