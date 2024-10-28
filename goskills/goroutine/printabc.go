package main

import (
	"fmt"
)

var printcount int = 5

func printLetter(ch chan bool, letter string, nextCh chan bool, done chan bool) {
	for i := 0; i < printcount; i++ {
		<-ch
		fmt.Println(letter)

		//当最后一个字母打印完最后一次的时候，需要通知其他几个协程退出，否则命中 死锁 循环等待的条件
		if letter == "c" && i == printcount-1 {
			done <- true // 通知其他协程任务已经完成
			break
		}
		nextCh <- true
	}
}

// 死锁条件：循环等待、互斥、非抢占、请求保持

/*
chan什么情况下会panic?
1、读写在一个协程
2、重复关闭chan
3、往一个已经关闭过的chan写数据
4、chan只有生产没有消费、反之
5、操作一个nil类型的chan
*/
func printABC() {
	ch1 := make(chan bool, 1)
	ch2 := make(chan bool, 1)
	ch3 := make(chan bool, 1)
	done := make(chan bool) // 用于通知任务完成的通道

	go printLetter(ch1, "a", ch2, done)
	go printLetter(ch2, "b", ch3, done)
	go printLetter(ch3, "c", ch1, done)

	ch1 <- true // 启动第一个协程

	<-done // 等待任务完成的通知
	close(ch1)
	close(ch2)
	close(ch3)

	//wg.Wait() // 等待最后一个协程完成
	close(done)
}

func main() {
	printABC()
}
