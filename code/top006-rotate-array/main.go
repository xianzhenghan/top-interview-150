package main

import "fmt"

// 最佳写法
func rotate(nums []int, k int) {
	copy(nums, append(nums[len(nums)-(k%len(nums)):], nums[:len(nums)-(k%len(nums))]...))
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}

func rotateV3(nums []int, k int) {
	for ; k > 0; k-- {
		copy(nums, append([]int{nums[len(nums)-1]}, nums[:len(nums)-1]...))
	}
}

func rotateV2(nums []int, k int) {
	if k == 0 {
		return
	}
	copy(nums, append([]int{nums[len(nums)-1]}, nums[:len(nums)-1]...))
	rotate(nums, k-1)
}
