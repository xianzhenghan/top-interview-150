package main

import (
	"fmt"
	_ "net/http/pprof"
	"sync"
)

func main2() {
	catChan := make(chan struct{}, 1)
	dogChan := make(chan struct{}, 1)
	fishChan := make(chan struct{}, 1)

	wg := &sync.WaitGroup{}

	wg.Add(3)
	catChan <- struct{}{}
	go printCat(wg, catChan, dogChan)
	go printDog(wg, dogChan, fishChan)
	go printFish(wg, fishChan, catChan)
	wg.Wait()
}

func printCat(wg *sync.WaitGroup, catChan chan struct{}, dogChan chan struct{}) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		select {
		case <-catChan:
			fmt.Println("Cat")
			dogChan <- struct{}{}
		}
	}
}

func printDog(wg *sync.WaitGroup, dogChan chan struct{}, fishChan chan struct{}) {
	defer wg.Done()

	for i := 0; i < 100; i++ {
		select {
		case <-dogChan:
			fmt.Println("Dog")
			fishChan <- struct{}{}
		}
	}

}

func printFish(wg *sync.WaitGroup, fishChan chan struct{}, catChan chan struct{}) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		select {
		case <-fishChan:
			fmt.Println("Fish")
			catChan <- struct{}{}
		}
	}

}
