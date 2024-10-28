package main

import "fmt"

func main() {
	test()
}

//
func test() {
	ch := make(chan struct{})
	close(ch)
	for {
		select {
		case <-ch:
			fmt.Println("ch")
		}
	}
}
