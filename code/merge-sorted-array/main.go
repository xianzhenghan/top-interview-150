package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{6, 7, 8, 9, 10}
	merge(nums1, len(nums1), nums2, len(nums2))
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

}
