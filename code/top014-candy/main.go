package main

import (
	"fmt"
	"sort"
	"sync"
)

//https://leetcode.cn/problems/candy/description/?envType=study-plan-v2&envId=top-interview-150
//n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。
//
//你需要按照以下要求，给这些孩子分发糖果：
//
//每个孩子至少分配到 1 个糖果。
//相邻两个孩子评分更高的孩子会获得更多的糖果。
//请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。
//
//
//
//示例 1：
//
//输入：ratings = [1,0,2]
//输出：5
//解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。
//示例 2：
//
//输入：ratings = [1,2,2]
//输出：4
//解释：你可以分别给第一个、第二个、第三个孩子分发 1、2、1 颗糖果。
//     第三个孩子只得到 1 颗糖果，这满足题面中的两个条件。
//
//
//提示：
//
//n == ratings.length
//1 <= n <= 2 * 104
//0 <= ratings[i] <= 2 * 104

func candy(ratings []int) int {
	sort.Ints(ratings)
	dp := make([]int, len(ratings))
	dp[0] = 1
	sum := 1
	fmt.Println(ratings)
	for i := 1; i < len(ratings); i++ {
		if ratings[i] == ratings[i-1] {
			dp[i] = dp[i-1]
			sum += dp[i-1]
		} else {
			dp[i] = dp[i-1] + 1
			sum += dp[i-1] + 1
		}
	}
	fmt.Println(dp)
	return sum
}

// 消费者非阻塞 channel 如果close 需要成产者消费者 退出
func consumer1(ch chan int, wg *sync.WaitGroup) {
	for {
		select {
		case v, ok := <-ch:
			fmt.Println("get from channel :", v, ok)
			if !ok {
				fmt.Println("channel os close")
				wg.Done()
				break
			}

		default:
		}
	}

}

//
func producer1(ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("producer:", i)
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go producer1(ch)
	go consumer1(ch, &wg)

	wg.Wait()
}
