package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	//非阻塞接收，有可以用就直接用，没有就默认，不等待
	fmt.Println("程序开始执行...")
	go func() {
		time.Sleep(time.Second * 2)
		messages <- "结果  num one"
	}()

	//如果存在default 那么就是非阻塞 有可以用就直接用，没有就默认，不等待。
	//如果将下面的default语句块注释，那么select会等待接收messages中的值。
	select {
	case msg := <-messages:
		fmt.Println("A 收到消息", msg)
	default:
		fmt.Println("C 没有消息收到")

	}

	msg := "hi baby"
	select {
	case messages <- msg:
		fmt.Println("B sent message", msg)
	default:
		fmt.Println("B no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("C收到消息", msg)
	case sig := <-signals:
		fmt.Println("C收到消息", sig)
	default:
		fmt.Println("C no  activity")
	}
}
